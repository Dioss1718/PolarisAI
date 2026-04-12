package server

import (
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/diya-suryawanshi/cloud/backend/orchestrator"
)

type gitopsApprovalRecord struct {
	ApprovalID    string  `json:"approvalId"`
	NodeID        string  `json:"nodeId"`
	Action        string  `json:"action"`
	Score         float64 `json:"score"`
	Reason        string  `json:"reason"`
	Status        string  `json:"status"`
	Message       string  `json:"message"`
	URL           string  `json:"url,omitempty"`
	PRNumber      int     `json:"prNumber,omitempty"`
	Branch        string  `json:"branch,omitempty"`
	RequestedAt   string  `json:"requestedAt"`
	ReviewedAt    string  `json:"reviewedAt,omitempty"`
	ReviewedBy    string  `json:"reviewedBy,omitempty"`
	ReviewComment string  `json:"reviewComment,omitempty"`
}

type gitopsAuditEvent struct {
	Timestamp     string `json:"timestamp"`
	Actor         string `json:"actor"`
	Action        string `json:"action"`
	ApprovalID    string `json:"approvalId"`
	NodeID        string `json:"nodeId"`
	FinalAction   string `json:"finalAction"`
	Status        string `json:"status"`
	ReviewComment string `json:"reviewComment,omitempty"`
}

type gitopsApprovalBucket struct {
	Items map[string]*gitopsApprovalRecord
	Audit []gitopsAuditEvent
}

type runtimeManager struct {
	mu           sync.Mutex
	lastScenario string
	lastSeed     int
	lastState    *orchestrator.PipelineResult
	buckets      map[string]*gitopsApprovalBucket
}

func newRuntimeManager() *runtimeManager {
	return &runtimeManager{
		buckets: make(map[string]*gitopsApprovalBucket),
	}
}

func runtimeKey(scenario string, seed int) string {
	return fmt.Sprintf("%s|%d", scenario, seed)
}

func (rm *runtimeManager) setLatest(result *orchestrator.PipelineResult, scenario string, seed int) {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	rm.lastScenario = scenario
	rm.lastSeed = seed
	rm.lastState = result
}

func (rm *runtimeManager) latestFor(scenario string, seed int) *orchestrator.PipelineResult {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	if rm.lastState == nil {
		return nil
	}
	if rm.lastScenario != scenario || rm.lastSeed != seed {
		return nil
	}
	return rm.lastState
}

func (rm *runtimeManager) seedApprovalsFromPipeline(result *orchestrator.PipelineResult, scenario string, seed int) {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	key := runtimeKey(scenario, seed)
	existing := rm.buckets[key]

	audit := []gitopsAuditEvent{}
	if existing != nil {
		audit = existing.Audit
	}

	bucket := &gitopsApprovalBucket{
		Items: make(map[string]*gitopsApprovalRecord),
		Audit: audit,
	}

	now := time.Now().UTC().Format(time.RFC3339)

	for _, rec := range result.Recommendations {
		if rec.Status == "REJECTED" {
			continue
		}

		approvalID := fmt.Sprintf("%s|%s", rec.NodeID, rec.FinalAction)

		bucket.Items[approvalID] = &gitopsApprovalRecord{
			ApprovalID:  approvalID,
			NodeID:      rec.NodeID,
			Action:      rec.FinalAction,
			Score:       rec.Score,
			Reason:      rec.Reason,
			Status:      "PENDING_APPROVAL",
			Message:     "Awaiting explicit human approval before PR creation",
			RequestedAt: now,
		}
	}

	rm.buckets[key] = bucket
}

func (rm *runtimeManager) getApproval(scenario string, seed int, approvalID string) (*gitopsApprovalRecord, bool) {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	bucket := rm.buckets[runtimeKey(scenario, seed)]
	if bucket == nil {
		return nil, false
	}

	item, ok := bucket.Items[approvalID]
	return item, ok
}

func (rm *runtimeManager) addAuditEvent(scenario string, seed int, event gitopsAuditEvent) {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	key := runtimeKey(scenario, seed)
	bucket := rm.buckets[key]
	if bucket == nil {
		bucket = &gitopsApprovalBucket{
			Items: map[string]*gitopsApprovalRecord{},
			Audit: []gitopsAuditEvent{},
		}
		rm.buckets[key] = bucket
	}

	bucket.Audit = append(bucket.Audit, event)
}

func (rm *runtimeManager) auditFor(scenario string, seed int) []gitopsAuditEvent {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	bucket := rm.buckets[runtimeKey(scenario, seed)]
	if bucket == nil {
		return []gitopsAuditEvent{}
	}

	out := make([]gitopsAuditEvent, len(bucket.Audit))
	copy(out, bucket.Audit)

	sort.Slice(out, func(i, j int) bool {
		return out[i].Timestamp > out[j].Timestamp
	})

	return out
}

func (rm *runtimeManager) overlayGitOps(result *orchestrator.PipelineResult, scenario string, seed int) *orchestrator.PipelineResult {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	if result == nil {
		return nil
	}

	cloned := *result
	cloned.GitOps = rm.buildGitOpsDTOLocked(scenario, seed)
	return &cloned
}

func (rm *runtimeManager) buildGitOpsDTOLocked(scenario string, seed int) orchestrator.GitOpsDTO {
	bucket := rm.buckets[runtimeKey(scenario, seed)]
	if bucket == nil || len(bucket.Items) == 0 {
		return orchestrator.GitOpsDTO{
			Status:  "idle",
			Message: "No GitOps approval requests available for this run",
			PRs:     []orchestrator.GitOpsPRDTO{},
		}
	}

	items := make([]*gitopsApprovalRecord, 0, len(bucket.Items))
	var pendingCount, approvedCount, rejectedCount, createdCount, failedCount int

	for _, item := range bucket.Items {
		items = append(items, item)

		switch item.Status {
		case "PENDING_APPROVAL":
			pendingCount++
		case "REJECTED":
			rejectedCount++
		case "APPROVED":
			approvedCount++
		case "PR_CREATED", "MERGED":
			createdCount++
		case "FAILED", "BLOCKED":
			failedCount++
		}
	}

	sort.Slice(items, func(i, j int) bool {
		if items[i].Status == items[j].Status {
			return items[i].NodeID < items[j].NodeID
		}
		return items[i].Status < items[j].Status
	})

	prs := make([]orchestrator.GitOpsPRDTO, 0, len(items))
	for _, item := range items {
		prs = append(prs, orchestrator.GitOpsPRDTO{
			ApprovalID:    item.ApprovalID,
			URL:           item.URL,
			Status:        item.Status,
			PRNumber:      item.PRNumber,
			Branch:        item.Branch,
			NodeID:        item.NodeID,
			Action:        item.Action,
			Message:       item.Message,
			RequestedAt:   item.RequestedAt,
			ReviewedAt:    item.ReviewedAt,
			ReviewedBy:    item.ReviewedBy,
			ReviewComment: item.ReviewComment,
		})
	}

	status := "pending_approval"
	message := fmt.Sprintf(
		"%d pending, %d rejected, %d approved, %d PR-created, %d blocked/failed",
		pendingCount,
		rejectedCount,
		approvedCount,
		createdCount,
		failedCount,
	)

	return orchestrator.GitOpsDTO{
		Status:  status,
		Message: message,
		PRs:     prs,
	}
}
