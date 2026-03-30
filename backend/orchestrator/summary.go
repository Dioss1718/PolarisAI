package orchestrator

import (
	"math"

	"github.com/diya-suryawanshi/cloud/carbon"
	"github.com/diya-suryawanshi/cloud/graph-engine/graph"
)

func complianceScoreFromGraph(g *graph.Graph, risks map[string]float64) float64 {
	score := 100.0

	for _, n := range g.Nodes {
		if n.Exposure == "PUBLIC" {
			score -= 7
		}
		if n.Type == "DATABASE" || n.Type == "IAM_ROLE" {
			score -= 4
		}
		score -= float64(len(n.Compliance)) * 1.4
		if risks[n.ID] >= 8 {
			score -= 3.5
		}
	}

	if score < 0 {
		score = 0
	}
	return math.Round(score*100) / 100
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
	return math.Round(score*100) / 100
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
	return math.Round((total/float64(len(g.Nodes)))*100) / 100
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
		AvgAttackPathLength: math.Round(attack.AvgPathLength*100) / 100,
		ReachableNodes:      attack.ReachableNodes,
		HighRiskCount:       highRisk,
		PublicExposureCount: publicExposure,
		ApprovedCount:       approved,
		ModifiedCount:       modified,
		RejectedCount:       rejected,
		UrgentCount:         urgent,
		BillShockCount:      billShock,
		CurrentTotalCost:    math.Round(totalCost*100) / 100,
		Forecast30Total:     math.Round(f30*100) / 100,
		Forecast90Total:     math.Round(f90*100) / 100,
		AverageRisk:         math.Round(avgRisk*100) / 100,
		TotalRisk:           math.Round(totalRisk*100) / 100,
		CurrentCarbonTotal:  math.Round(carbonReport.Total*100) / 100,
		ComplianceScore:     complianceScoreFromGraph(g, risks),
		CostRiskScore:       costRiskFromForecasts(forecasts),
	}
}

func BuildProjectedSummary(
	g *graph.Graph,
	projectedAttack AttackMetrics,
	projectedRisks map[string]float64,
	projectedCarbon carbon.Report,
	currentCarbon float64,
	currentAverageRisk float64,
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

	reductionPct := carbon.Compare(currentCarbon, projectedCarbon.Total)
	greenScore := carbon.GreenScore(currentCarbon, projectedCarbon.Total)

	riskReductionPct := 0.0
	if currentAverageRisk > 0 {
		riskReductionPct = ((currentAverageRisk - avgRisk) / currentAverageRisk) * 100
		if riskReductionPct < 0 {
			riskReductionPct = 0
		}
	}

	return ProjectedSummaryDTO{
		ProjectedTotalCost:        math.Round(projectedCost*100) / 100,
		ProjectedAttackPathCount:  projectedAttack.PathCount,
		ProjectedPublicExposure:   publicExposure,
		ProjectedAverageRisk:      math.Round(avgRisk*100) / 100,
		ProjectedCarbonTotal:      math.Round(projectedCarbon.Total*100) / 100,
		CarbonReductionPct:        math.Round(reductionPct*100) / 100,
		GreenScore:                math.Round(greenScore*100) / 100,
		ProjectedComplianceScore:  complianceScoreFromGraph(g, projectedRisks),
		ProjectedCostRiskScore:    projectedCostRiskFromGraph(g, projectedRisks),
		ProjectedRiskReductionPct: math.Round(riskReductionPct*100) / 100,
	}
}
