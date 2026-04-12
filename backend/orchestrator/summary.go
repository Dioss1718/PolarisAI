package orchestrator

import (
	"math"
	"sort"

	"github.com/diya-suryawanshi/cloud/carbon"
	"github.com/diya-suryawanshi/cloud/graph-engine/graph"
)

func complianceScoreFromGraph(g *graph.Graph, risks map[string]float64) float64 {
	score := 100.0

	negativeFlags := map[string]bool{
		"ADMIN_ACCESS":       true,
		"IAM_OVERPRIVILEGED": true,
		"OPEN_PORT":          true,
		"PORT_0_0_0_0":       true,
		"PUBLIC_BUCKET":      true,
	}

	for _, n := range g.Nodes {
		if n.Exposure == "PUBLIC" {
			score -= 7
		}
		if n.Type == "DATABASE" || n.Type == "IAM_ROLE" {
			score -= 4
		}

		negativeCount := 0
		for _, flag := range n.Compliance {
			if negativeFlags[flag] {
				negativeCount++
			}
		}
		score -= float64(negativeCount) * 1.4

		if risks[n.ID] >= 8 {
			score -= 3.5
		}
	}

	if score < 0 {
		score = 0
	}
	return roundSummary2(score)
}

func costRiskFromForecasts(forecasts []ForecastDTO) float64 {
	if len(forecasts) == 0 {
		return 0
	}

	total := 0.0
	for _, f := range forecasts {
		if f.CurrentCost > 0 {
			total += ((f.Forecast90 - f.CurrentCost) / f.CurrentCost) * 100
		}
		if f.BillShock {
			total += 8
		}
	}

	score := total / float64(len(forecasts))
	if score < 0 {
		score = 0
	}
	return roundSummary2(score)
}

func projectedCostRiskFromGraph(g *graph.Graph, risks map[string]float64) float64 {
	total := 0.0
	for _, n := range g.Nodes {
		total += (n.Cost / 20.0) + risks[n.ID]*1.8
		if n.Exposure == "PUBLIC" {
			total += 6
		}
	}
	if len(g.Nodes) == 0 {
		return 0
	}
	return roundSummary2(total / float64(len(g.Nodes)))
}

func buildTopCarbonSources(report carbon.Report) []CarbonSourceDTO {
	if report.Total <= 0 || len(report.Top) == 0 {
		return []CarbonSourceDTO{}
	}

	out := make([]CarbonSourceDTO, 0, len(report.Top))
	for _, item := range report.Top {
		pct := 0.0
		if report.Total > 0 {
			pct = (item.Value / report.Total) * 100
		}

		out = append(out, CarbonSourceDTO{
			NodeID:              item.NodeID,
			Carbon:              roundSummary2(item.Value),
			PercentContribution: roundSummary2(pct),
		})
	}
	return out
}

func buildCarbonActionImpact(
	currentReport carbon.Report,
	projectedReport carbon.Report,
	recommendations []RecommendationDTO,
) []CarbonActionImpactDTO {
	currentByNode := make(map[string]float64, len(currentReport.Results))
	for _, item := range currentReport.Results {
		currentByNode[item.NodeID] = item.Value
	}

	projectedByNode := make(map[string]float64, len(projectedReport.Results))
	for _, item := range projectedReport.Results {
		projectedByNode[item.NodeID] = item.Value
	}

	selectedByNode := make(map[string]RecommendationDTO)
	for _, rec := range recommendations {
		if rec.Status == "APPROVED" || rec.Status == "MODIFIED" {
			selectedByNode[rec.NodeID] = rec
		}
	}

	impacts := make([]CarbonActionImpactDTO, 0, len(selectedByNode))
	for nodeID, rec := range selectedByNode {
		before := currentByNode[nodeID]
		after := projectedByNode[nodeID]
		reduction := before - after

		if before == 0 && after == 0 {
			continue
		}

		impacts = append(impacts, CarbonActionImpactDTO{
			NodeID:          nodeID,
			Action:          rec.FinalAction,
			CarbonBefore:    roundSummary2(before),
			CarbonAfter:     roundSummary2(after),
			CarbonReduction: roundSummary2(reduction),
		})
	}

	sort.Slice(impacts, func(i, j int) bool {
		return impacts[i].CarbonReduction > impacts[j].CarbonReduction
	})

	if len(impacts) > 5 {
		impacts = impacts[:5]
	}

	return impacts
}

func BuildCurrentSummary(
	g *graph.Graph,
	attack AttackMetrics,
	recommendations []RecommendationDTO,
	forecasts []ForecastDTO,
	risks map[string]float64,
	carbonReport carbon.Report,
) SummaryDTO {
	var (
		highRisk, publicExposure             int
		approved, modified, rejected, urgent int
		totalCost, f30, f90, totalRisk       float64
		billShock                            int
	)

	for _, n := range g.Nodes {
		totalCost += n.Cost
		if n.Exposure == "PUBLIC" {
			publicExposure++
		}
		if risks[n.ID] >= 8 {
			highRisk++
		}
		totalRisk += risks[n.ID]
	}

	for _, r := range recommendations {
		switch r.Status {
		case "APPROVED":
			approved++
		case "MODIFIED":
			modified++
		case "REJECTED":
			rejected++
		}
		if (r.Status == "APPROVED" || r.Status == "MODIFIED") && r.Risk >= 7 {
			urgent++
		}
	}

	for _, f := range forecasts {
		f30 += f.Forecast30
		f90 += f.Forecast90
		if f.BillShock {
			billShock++
		}
	}

	avgRisk := 0.0
	if len(g.Nodes) > 0 {
		avgRisk = totalRisk / float64(len(g.Nodes))
	}

	return SummaryDTO{
		TotalNodes:          len(g.Nodes),
		TotalEdges:          len(g.Edges),
		AttackPathCount:     attack.PathCount,
		AvgAttackPathLength: roundSummary2(attack.AvgPathLength),
		ReachableNodes:      attack.ReachableNodes,
		HighRiskCount:       highRisk,
		PublicExposureCount: publicExposure,
		ApprovedCount:       approved,
		ModifiedCount:       modified,
		RejectedCount:       rejected,
		UrgentCount:         urgent,
		BillShockCount:      billShock,
		CurrentTotalCost:    roundSummary2(totalCost),
		Forecast30Total:     roundSummary2(f30),
		Forecast90Total:     roundSummary2(f90),
		AverageRisk:         roundSummary2(avgRisk),
		TotalRisk:           roundSummary2(totalRisk),
		CurrentCarbonTotal:  roundSummary2(carbonReport.Total),
		ComplianceScore:     complianceScoreFromGraph(g, risks),
		CostRiskScore:       costRiskFromForecasts(forecasts),
		TopCarbonSources:    buildTopCarbonSources(carbonReport),
	}
}

func BuildProjectedSummary(
	g *graph.Graph,
	projectedAttack AttackMetrics,
	projectedRisks map[string]float64,
	currentCarbon carbon.Report,
	projectedCarbon carbon.Report,
	currentAverageRisk float64,
	recommendations []RecommendationDTO,
) ProjectedSummaryDTO {
	var projectedCost, totalRisk float64
	var publicExposure int

	for _, n := range g.Nodes {
		projectedCost += n.Cost
		if n.Exposure == "PUBLIC" {
			publicExposure++
		}
		totalRisk += projectedRisks[n.ID]
	}

	avgRisk := 0.0
	if len(g.Nodes) > 0 {
		avgRisk = totalRisk / float64(len(g.Nodes))
	}

	reductionPct := carbon.Compare(currentCarbon.Total, projectedCarbon.Total)
	greenScore := carbon.GreenScore(currentCarbon.Total, projectedCarbon.Total)

	riskReductionPct := 0.0
	if currentAverageRisk > 0 {
		riskReductionPct = ((currentAverageRisk - avgRisk) / currentAverageRisk) * 100
		if riskReductionPct < 0 {
			riskReductionPct = 0
		}
	}

	return ProjectedSummaryDTO{
		ProjectedTotalCost:        roundSummary2(projectedCost),
		ProjectedAttackPathCount:  projectedAttack.PathCount,
		ProjectedPublicExposure:   publicExposure,
		ProjectedAverageRisk:      roundSummary2(avgRisk),
		ProjectedCarbonTotal:      roundSummary2(projectedCarbon.Total),
		CarbonReductionPct:        roundSummary2(reductionPct),
		GreenScore:                roundSummary2(greenScore),
		ProjectedComplianceScore:  complianceScoreFromGraph(g, projectedRisks),
		ProjectedCostRiskScore:    projectedCostRiskFromGraph(g, projectedRisks),
		ProjectedRiskReductionPct: roundSummary2(riskReductionPct),
		TopCarbonSources:          buildTopCarbonSources(projectedCarbon),
		CarbonActionImpact:        buildCarbonActionImpact(currentCarbon, projectedCarbon, recommendations),
	}
}

func roundSummary2(v float64) float64 {
	return math.Round(v*100) / 100
}
