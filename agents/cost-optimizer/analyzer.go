package costoptimizer

import "github.com/diya-suryawanshi/cloud/graph-engine/graph"

func AnalyzeCostSignals(g *graph.Graph) []CostSignal {

	var signals []CostSignal

	for _, node := range g.Nodes {

		if node.Cost == 0 {
			continue
		}

		graphImpact := ComputeGraphImpact(g, node.ID)

		waste := node.Cost * (1 - node.Utilization/100)

		forecast := ForecastCost(node.Cost, node.Utilization)

		confidence := 1.0 - (graphImpact * 0.05)
		if confidence < 0.3 {
			confidence = 0.3
		}

		signals = append(signals, CostSignal{
			NodeID:                node.ID,
			ResourceType:          node.Type,
			CurrentCost:           node.Cost,
			Utilization:           node.Utilization,
			WasteScore:            waste,
			OptimizationPotential: waste / node.Cost,
			GraphImpact:           graphImpact,
			ForecastCost:          forecast,
			Confidence:            confidence,
			Reason:                "Graph-aware cost inefficiency",
		})
	}

	return signals
}
