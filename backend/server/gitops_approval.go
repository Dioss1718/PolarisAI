package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/diya-suryawanshi/cloud/backend/auth"
	"github.com/diya-suryawanshi/cloud/backend/orchestrator"
	pluginpkg "github.com/diya-suryawanshi/cloud/backend/plugin"
	"github.com/diya-suryawanshi/cloud/gitops"
	graphpkg "github.com/diya-suryawanshi/cloud/graph-engine/graph"
	modelspkg "github.com/diya-suryawanshi/cloud/graph-engine/models"
	"github.com/diya-suryawanshi/cloud/rbac"
)

type gitopsReviewRequest struct {
	Scenario   string `json:"scenario"`
	Seed       int    `json:"seed"`
	ApprovalID string `json:"approvalId"`
	Comment    string `json:"comment"`
}

func canReviewGitOps(session auth.Session) bool {
	return session.Features[string(rbac.FeatureGitOpsMerge)] == string(rbac.AccessFull)
}

func handleGitOpsApprove(w http.ResponseWriter, r *http.Request, session auth.Session) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	if !canReviewGitOps(session) {
		writeJSON(w, http.StatusForbidden, map[string]string{"error": "role not allowed to approve GitOps actions"})
		return
	}

	var req gitopsReviewRequest
	if err := decodeJSONBody(r, &req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	req.Scenario = strings.TrimSpace(req.Scenario)
	if req.Scenario == "" {
		req.Scenario = "FULL_CHAOS"
	}
	if req.Seed == 0 {
		req.Seed = 42
	}
	req.ApprovalID = strings.TrimSpace(req.ApprovalID)
	if req.ApprovalID == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "approvalId is required"})
		return
	}

	result := runtimeState.latestFor(req.Scenario, req.Seed)
	if result == nil {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "no pipeline state found for the requested scenario and seed"})
		return
	}

	record, ok := runtimeState.getApproval(req.Scenario, req.Seed, req.ApprovalID)
	if !ok {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "approval request not found"})
		return
	}

	if record.Status != "PENDING_APPROVAL" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "approval request is no longer pending"})
		return
	}

	if pluginpkg.GitOps == nil {
		record.Status = "FAILED"
		record.Message = "GitOps plugin not configured"
		record.ReviewedAt = time.Now().UTC().Format(time.RFC3339)
		record.ReviewedBy = session.EmployeeID
		record.ReviewComment = strings.TrimSpace(req.Comment)

		runtimeState.addAuditEvent(req.Scenario, req.Seed, gitopsAuditEvent{
			Timestamp:     time.Now().UTC().Format(time.RFC3339),
			Actor:         session.EmployeeID,
			Action:        "APPROVE_ATTEMPT",
			ApprovalID:    record.ApprovalID,
			NodeID:        record.NodeID,
			FinalAction:   record.Action,
			Status:        record.Status,
			ReviewComment: record.ReviewComment,
		})

		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "GitOps plugin not configured"})
		return
	}

	currentGraph := graphFromPipelineResult(result)
	nodeRisks := nodeRiskMapFromResult(result)

	prs, err := pluginpkg.GitOps.Run(currentGraph, []gitops.Decision{
		{
			NodeID:      record.NodeID,
			Action:      record.Action,
			FinalAction: record.Action,
			Score:       record.Score,
			Reason:      record.Reason,
		},
	}, nodeRisks)

	record.ReviewedAt = time.Now().UTC().Format(time.RFC3339)
	record.ReviewedBy = session.EmployeeID
	record.ReviewComment = strings.TrimSpace(req.Comment)

	if err != nil {
		record.Status = "FAILED"
		record.Message = err.Error()
	} else if len(prs) == 0 {
		record.Status = "FAILED"
		record.Message = "GitOps returned no PR response"
	} else {
		pr := prs[0]
		record.Status = normalizeApprovalResultStatus(pr.Status)
		record.Message = pr.Message
		record.URL = pr.URL
		record.PRNumber = pr.PRNumber
		record.Branch = pr.Branch
	}

	runtimeState.addAuditEvent(req.Scenario, req.Seed, gitopsAuditEvent{
		Timestamp:     time.Now().UTC().Format(time.RFC3339),
		Actor:         session.EmployeeID,
		Action:        "APPROVE",
		ApprovalID:    record.ApprovalID,
		NodeID:        record.NodeID,
		FinalAction:   record.Action,
		Status:        record.Status,
		ReviewComment: record.ReviewComment,
	})

	updated := runtimeState.overlayGitOps(result, req.Scenario, req.Seed)
	writeJSON(w, http.StatusOK, updated.GitOps)
}

func handleGitOpsReject(w http.ResponseWriter, r *http.Request, session auth.Session) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	if !canReviewGitOps(session) {
		writeJSON(w, http.StatusForbidden, map[string]string{"error": "role not allowed to reject GitOps actions"})
		return
	}

	var req gitopsReviewRequest
	if err := decodeJSONBody(r, &req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	req.Scenario = strings.TrimSpace(req.Scenario)
	if req.Scenario == "" {
		req.Scenario = "FULL_CHAOS"
	}
	if req.Seed == 0 {
		req.Seed = 42
	}
	req.ApprovalID = strings.TrimSpace(req.ApprovalID)
	if req.ApprovalID == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "approvalId is required"})
		return
	}

	result := runtimeState.latestFor(req.Scenario, req.Seed)
	if result == nil {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "no pipeline state found for the requested scenario and seed"})
		return
	}

	record, ok := runtimeState.getApproval(req.Scenario, req.Seed, req.ApprovalID)
	if !ok {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "approval request not found"})
		return
	}

	if record.Status != "PENDING_APPROVAL" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "approval request is no longer pending"})
		return
	}

	record.Status = "REJECTED"
	record.Message = "Rejected by reviewer"
	record.ReviewedAt = time.Now().UTC().Format(time.RFC3339)
	record.ReviewedBy = session.EmployeeID
	record.ReviewComment = strings.TrimSpace(req.Comment)

	runtimeState.addAuditEvent(req.Scenario, req.Seed, gitopsAuditEvent{
		Timestamp:     time.Now().UTC().Format(time.RFC3339),
		Actor:         session.EmployeeID,
		Action:        "REJECT",
		ApprovalID:    record.ApprovalID,
		NodeID:        record.NodeID,
		FinalAction:   record.Action,
		Status:        record.Status,
		ReviewComment: record.ReviewComment,
	})

	updated := runtimeState.overlayGitOps(result, req.Scenario, req.Seed)
	writeJSON(w, http.StatusOK, updated.GitOps)
}

func handleGitOpsAudit(w http.ResponseWriter, r *http.Request, session auth.Session) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	scenario := strings.TrimSpace(r.URL.Query().Get("scenario"))
	if scenario == "" {
		scenario = "FULL_CHAOS"
	}

	seed := 42
	if s := strings.TrimSpace(r.URL.Query().Get("seed")); s != "" {
		var parsed int
		_, _ = fmt.Sscanf(s, "%d", &parsed)
		if parsed != 0 {
			seed = parsed
		}
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"items": runtimeState.auditFor(scenario, seed),
	})
}

func handleGitOpsRefreshPR(w http.ResponseWriter, r *http.Request, session auth.Session) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	var req gitopsReviewRequest
	if err := decodeJSONBody(r, &req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	req.Scenario = strings.TrimSpace(req.Scenario)
	if req.Scenario == "" {
		req.Scenario = "FULL_CHAOS"
	}
	if req.Seed == 0 {
		req.Seed = 42
	}
	req.ApprovalID = strings.TrimSpace(req.ApprovalID)
	if req.ApprovalID == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "approvalId is required"})
		return
	}

	result := runtimeState.latestFor(req.Scenario, req.Seed)
	if result == nil {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "no pipeline state found for the requested scenario and seed"})
		return
	}

	record, ok := runtimeState.getApproval(req.Scenario, req.Seed, req.ApprovalID)
	if !ok {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "approval request not found"})
		return
	}

	if record.PRNumber == 0 {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "no PR exists for this approval request yet"})
		return
	}

	state, merged, err := gitops.GetPRStatus(record.PRNumber)
	if err != nil {
		writeJSON(w, http.StatusBadGateway, map[string]string{"error": err.Error()})
		return
	}

	if merged {
		record.Status = "MERGED"
		record.Message = "PR merged into main"
	} else {
		record.Message = "PR state refreshed: " + state
	}

	runtimeState.addAuditEvent(req.Scenario, req.Seed, gitopsAuditEvent{
		Timestamp:   time.Now().UTC().Format(time.RFC3339),
		Actor:       session.EmployeeID,
		Action:      "REFRESH_PR",
		ApprovalID:  record.ApprovalID,
		NodeID:      record.NodeID,
		FinalAction: record.Action,
		Status:      record.Status,
	})

	updated := runtimeState.overlayGitOps(result, req.Scenario, req.Seed)
	writeJSON(w, http.StatusOK, updated.GitOps)
}

func decodeJSONBody(r *http.Request, dst interface{}) error {
	return jsonNewDecoder(r, dst)
}

func jsonNewDecoder(r *http.Request, dst interface{}) error {
	return json.NewDecoder(r.Body).Decode(dst)
}

func normalizeApprovalResultStatus(status string) string {
	status = strings.ToUpper(strings.TrimSpace(status))
	switch status {
	case "CREATED":
		return "PR_CREATED"
	case "BLOCKED":
		return "BLOCKED"
	case "FAILED":
		return "FAILED"
	default:
		return status
	}
}

func graphFromPipelineResult(result *orchestrator.PipelineResult) *graphpkg.Graph {
	g := graphpkg.NewGraph()

	for _, n := range result.Nodes {
		g.AddNode(modelspkg.Node{
			ID:          n.ID,
			Type:        n.Type,
			Name:        n.Label,
			Cloud:       n.Cloud,
			Region:      n.Region,
			Environment: n.Environment,
			Cost:        n.Cost,
			Utilization: n.Utilization,
			Exposure:    n.Exposure,
			Criticality: n.Criticality,
			Compliance:  append([]string{}, n.Compliance...),
		})
	}

	for _, e := range result.Edges {
		g.AddEdge(modelspkg.Edge{
			From:   e.From,
			To:     e.To,
			Type:   e.Type,
			Weight: e.Weight,
		})
	}

	return g
}

func nodeRiskMapFromResult(result *orchestrator.PipelineResult) map[string]float64 {
	out := make(map[string]float64, len(result.Nodes))
	for _, n := range result.Nodes {
		out[n.ID] = n.Risk
	}
	return out
}
