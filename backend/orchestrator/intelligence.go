package orchestrator

import (
	"fmt"
	"math"
	"sort"
	"strings"

	actiongenerator "github.com/diya-suryawanshi/cloud/agents/action-generator"
	"github.com/diya-suryawanshi/cloud/graph-engine/graph"
)

func round2(v float64) float64 {
	return math.Round(v*100) / 100
}

func clamp(v, lo, hi float64) float64 {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}

func computeBlastRadius(g *graph.Graph) map[string]int {
	result := make(map[string]int)

	for id := range g.Nodes {
		seen := map[string]bool{id: true}
		queue := []string{id}

		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]

			for _, e := range g.Adjacency[cur] {
				if !seen[e.To] {
					seen[e.To] = true
					queue = append(queue, e.To)
				}
			}
		}
		result[id] = len(seen) - 1
	}

	return result
}

func attackMembership(attackPaths [][]string) map[string]int {
	m := make(map[string]int)
	for _, path := range attackPaths {
		seen := map[string]bool{}
		for _, n := range path {
			if !seen[n] {
				m[n]++
				seen[n] = true
			}
		}
	}
	return m
}

func buildReasonForNode(g *graph.Graph, nodeID string, nodeRisk float64, attackPaths [][]string) string {
	node := g.Nodes[nodeID]
	if node.ID == "" {
		return "Node is part of the governance graph."
	}

	for _, path := range attackPaths {
		for i, p := range path {
			if p == nodeID {
				prev := ""
				next := ""
				if i > 0 {
					prev = path[i-1]
				}
				if i < len(path)-1 {
					next = path[i+1]
				}

				parts := []string{}
				if node.Exposure == "PUBLIC" {
					parts = append(parts, "it is publicly exposed")
				}
				if nodeRisk >= 8 {
					parts = append(parts, "its computed risk is high")
				}
				if prev != "" && next != "" {
					parts = append(parts, fmt.Sprintf("it connects %s → %s", prev, next))
				}

				if len(parts) > 0 {
					return "This node is risky because " + strings.Join(parts, ", ") + "."
				}
			}
		}
	}

	if node.Exposure == "PUBLIC" {
		return "This node is risky because it is publicly exposed and expands the reachable attack surface."
	}

	return "This node contributes to downstream operational and security risk in the cloud graph."
}

func buildNodeIntel(
	g *graph.Graph,
	risks map[string]float64,
	attackPaths [][]string,
	forecasts []ForecastDTO,
) []NodeIntelDTO {
	blastMap := computeBlastRadius(g)
	pathHits := attackMembership(attackPaths)

	forecastMap := make(map[string]ForecastDTO)
	for _, f := range forecasts {
		forecastMap[f.NodeID] = f
	}

	out := make([]NodeIntelDTO, 0, len(g.Nodes))
	for _, n := range g.Nodes {
		f := forecastMap[n.ID]
		costRisk := 0.0
		if f.CurrentCost > 0 && f.Forecast90 > 0 {
			costRisk = ((f.Forecast90 - f.CurrentCost) / f.CurrentCost) * 100.0
			if costRisk < 0 {
				costRisk = 0
			}
		}

		complianceBurden := 0.0
		if n.Exposure == "PUBLIC" {
			complianceBurden += 20
		}
		complianceBurden += float64(len(n.Compliance)) * 6
		if n.Type == "DATABASE" || n.Type == "IAM_ROLE" {
			complianceBurden += 12
		}

		riskInfluence := clamp(risks[n.ID]+float64(blastMap[n.ID])*0.35+float64(pathHits[n.ID])*0.9, 0, 10)

		out = append(out, NodeIntelDTO{
			NodeID:           n.ID,
			BlastRadius:      blastMap[n.ID],
			RiskInfluence:    round2(riskInfluence),
			AttackPathCount:  pathHits[n.ID],
			Exposed:          n.Exposure == "PUBLIC",
			Why:              buildReasonForNode(g, n.ID, risks[n.ID], attackPaths),
			CostRisk:         round2(costRisk),
			ComplianceBurden: round2(complianceBurden),
			AffectedNodes:    buildBlastAffectedNodes(g, n.ID, 8),
		})
	}

	sort.Slice(out, func(i, j int) bool {
		if out[i].BlastRadius == out[j].BlastRadius {
			return out[i].RiskInfluence > out[j].RiskInfluence
		}
		return out[i].BlastRadius > out[j].BlastRadius
	})

	return out
}

func deriveSafetyLevel(env, exposure string, score float64) string {
	if env == "PROD" && exposure == "PUBLIC" && score < 0.7 {
		return "Guarded"
	}
	if score >= 0.82 {
		return "High"
	}
	if score >= 0.62 {
		return "Moderate"
	}
	return "Low"
}

func deriveConfidence(actionScore, finalScore float64, env string) float64 {
	base := 0.55 + (actionScore * 0.18) + (finalScore * 0.22)
	if env == "PROD" {
		base -= 0.08
	}
	return round2(clamp(base, 0.35, 0.99))
}

func buildNegotiationTraces(
	actions []actiongenerator.Action,
	recommendations []RecommendationDTO,
) []NegotiationTraceDTO {
	grouped := make(map[string][]actiongenerator.Action)
	for _, a := range actions {
		grouped[a.NodeID] = append(grouped[a.NodeID], a)
	}

	var out []NegotiationTraceDTO

	for _, rec := range recommendations {
		cands := grouped[rec.NodeID]
		sort.Slice(cands, func(i, j int) bool {
			return cands[i].Score > cands[j].Score
		})

		var alts []NegotiationAlternativeDTO
		for _, c := range cands {
			actionName := c.ActionType + "_" + c.Variant
			if actionName == rec.Action || actionName == rec.FinalAction {
				continue
			}
			alts = append(alts, NegotiationAlternativeDTO{
				Action:        actionName,
				Score:         round2(c.Score),
				CostDelta:     round2(c.CostDelta),
				RiskReduction: round2(c.RiskReduction),
				Disruption:    round2(c.Disruption),
				Reason:        "Rejected because it delivered a weaker policy-safe tradeoff than the selected remediation.",
			})
			if len(alts) == 3 {
				break
			}
		}

		why := "Selected because it produced the strongest risk-cost tradeoff after policy validation."
		if rec.Status == "MODIFIED" {
			why = "Selected in modified form because the original candidate required a safer policy-constrained execution path."
		}
		if rec.Status == "REJECTED" {
			why = "No candidate satisfied the required policy and safety thresholds for execution."
		}

		out = append(out, NegotiationTraceDTO{
			NodeID:         rec.NodeID,
			SelectedAction: rec.FinalAction,
			SelectedScore:  round2(rec.Score),
			WhySelected:    why,
			Alternatives:   alts,
		})
	}

	return out
}

func buildStructuredAlerts(summary SummaryDTO, projected ProjectedSummaryDTO, recs []RecommendationDTO, intel []NodeIntelDTO) []AlertDTO {
	var alerts []AlertDTO

	if summary.AttackPathCount > 0 {
		alerts = append(alerts, AlertDTO{
			Severity:     "high",
			Title:        "Attack paths remain reachable",
			Metric:       fmt.Sprintf("%d paths", summary.AttackPathCount),
			Reason:       "Reachable attacker chains still exist between exposed assets and sensitive resources.",
			WorkspaceTab: "attackpaths",
		})
	}

	if summary.BillShockCount > 0 {
		alerts = append(alerts, AlertDTO{
			Severity:     "high",
			Title:        "Bill shock watch active",
			Metric:       fmt.Sprintf("%d flagged nodes", summary.BillShockCount),
			Reason:       "Forecasting engine detected rising cost pressure and possible future overspend.",
			WorkspaceTab: "billshock",
		})
	}

	if summary.ComplianceScore < 70 {
		alerts = append(alerts, AlertDTO{
			Severity:     "medium",
			Title:        "Compliance posture below target",
			Metric:       fmt.Sprintf("%.2f", summary.ComplianceScore),
			Reason:       "Policy posture indicates unresolved public exposure, sensitive dependencies, or drift risk.",
			WorkspaceTab: "compliance",
		})
	}

	if projected.ProjectedComplianceScore > summary.ComplianceScore {
		alerts = append(alerts, AlertDTO{
			Severity:     "info",
			Title:        "Projected compliance improvement available",
			Metric:       fmt.Sprintf("+%.2f", projected.ProjectedComplianceScore-summary.ComplianceScore),
			Reason:       "Approved governance actions improve policy posture in the projected state.",
			WorkspaceTab: "compliance",
		})
	}

	if len(intel) > 0 && intel[0].BlastRadius > 0 {
		alerts = append(alerts, AlertDTO{
			Severity:     "critical",
			Title:        "High blast radius node detected",
			Metric:       fmt.Sprintf("%s · %d affected", intel[0].NodeID, intel[0].BlastRadius),
			Reason:       "A highly connected resource can propagate risk across multiple downstream assets.",
			NodeID:       intel[0].NodeID,
			WorkspaceTab: "compliance",
		})
	}

	for _, r := range recs {
		if r.Status == "MODIFIED" && r.Risk >= 8 {
			alerts = append(alerts, AlertDTO{
				Severity:     "critical",
				Title:        "High-risk remediation constrained by policy",
				Metric:       r.NodeID,
				Reason:       "A high-risk node required a safer modified action instead of the original candidate.",
				NodeID:       r.NodeID,
				Action:       r.FinalAction,
				WorkspaceTab: "governance",
			})
			break
		}
	}

	return alerts
}
