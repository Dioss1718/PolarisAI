package costoptimizer

import "github.com/diya-suryawanshi/cloud/graph-engine/graph"

func Analyze(g *graph.Graph) []CostSignal {
	var signals []CostSignal

	for _, node := range g.Nodes {
		if node.Cost <= 0 {
			continue
		}

		wasteRatio := 1 - (node.Utilization / 100.0)
		graphImpact := ComputeGraphImpact(g, node.ID)
		score := ComputeScore(node.Cost, wasteRatio, graphImpact, node.Environment, node.Type, node.Exposure)
		confidence := ComputeConfidence(graphImpact, node.Environment, node.Type, wasteRatio)
		reason := BuildReason(node.Type, node.Environment, node.Exposure, node.Cost, node.Utilization, graphImpact)

		signals = append(signals, CostSignal{
			NodeID:      node.ID,
			Type:        node.Type,
			Cost:        node.Cost,
			Utilization: node.Utilization,
			WasteRatio:  wasteRatio,
			GraphImpact: graphImpact,
			Env:         node.Environment,
			Score:       score,
			Confidence:  confidence,
			Reason:      reason,
		})
	}

	return signals
}
