package costoptimizer

import (
	"github.com/diya-suryawanshi/cloud/graph-engine/graph"
)

const (
	LowUtilThreshold = 30.0
)

func AnalyzeCosts(g *graph.Graph) []CostInsight {
	var insights []CostInsight

	for _, node := range g.Nodes {

		// Skip zero-cost resources
		if node.Cost == 0 {
			continue
		}

		waste := 0.0
		reason := ""

		// Idle resource detection
		if node.Utilization < LowUtilThreshold {
			waste = node.Cost * (1 - node.Utilization/100)
			reason = "Low utilization resource"
		}

		// Overprovision detection (simple heuristic)
		if node.Utilization < 20 {
			waste += node.Cost * 0.2
			reason = "Overprovisioned resource"
		}

		if waste > 0 {
			insights = append(insights, CostInsight{
				NodeID:      node.ID,
				CurrentCost: node.Cost,
				Waste:       waste,
				Reason:      reason,
			})
		}
	}

	return insights
}
