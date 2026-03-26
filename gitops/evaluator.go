package gitops

type GraphMetrics struct {
	TotalRisk float64
	NodeCount int
}

func EvaluateGraph(g *Graph, nodeRisks map[string]float64) GraphMetrics {
	var totalRisk float64

	for id := range g.Nodes {
		totalRisk += nodeRisks[id]
	}

	return GraphMetrics{
		TotalRisk: totalRisk,
		NodeCount: len(g.Nodes),
	}
}

func SelectBestGraph(current *Graph, proposed *Graph, nodeRisks map[string]float64) *Graph {
	currentMetrics := EvaluateGraph(current, nodeRisks)
	proposedMetrics := EvaluateGraph(proposed, nodeRisks)

	if proposedMetrics.TotalRisk < currentMetrics.TotalRisk {
		return proposed
	}

	return current
}
