package gitops

// This struct stores basic metrics of a graph
type GraphMetrics struct {
	TotalRisk float64
	NodeCount int
}

// This function calculates total risk and node count of a graph
func EvaluateGraph(g *Graph, nodeRisks map[string]float64) GraphMetrics {

	var totalRisk float64

	// Loop through all nodes and sum their risk values
	for id := range g.Nodes {
		totalRisk += nodeRisks[id]
	}

	// Return calculated metrics
	return GraphMetrics{
		TotalRisk: totalRisk,
		NodeCount: len(g.Nodes),
	}
}

// This function compares current graph and proposed graph
// and selects the better one based on total risk
func SelectBestGraph(current *Graph, proposed *Graph, nodeRisks map[string]float64) *Graph {

	// Calculate metrics for both graphs
	currentMetrics := EvaluateGraph(current, nodeRisks)
	proposedMetrics := EvaluateGraph(proposed, nodeRisks)

	// Decision logic: graph with lower risk is better
	if proposedMetrics.TotalRisk < currentMetrics.TotalRisk {
		return proposed
	}

	// If no improvement, keep current graph
	return current
}
