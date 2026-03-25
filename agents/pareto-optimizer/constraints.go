package paretooptimizer

import "github.com/diya-suryawanshi/cloud/graph-engine/graph"

// Hard constraints (invalid actions)
func IsValid(g *graph.Graph, action Action) bool {
	node := g.Nodes[action.NodeID]

	// 🚫 Do not terminate PROD resources
	if node.Environment == "PROD" && action.ActionType == "TERMINATE" {
		return false
	}

	return true
}

// Soft penalty (trade-off cost)
func ConstraintPenalty(g *graph.Graph, action Action) float64 {
	node := g.Nodes[action.NodeID]

	degree := float64(len(g.Adjacency[action.NodeID]))
	centrality := degree / 5.0

	criticality := float64(node.Criticality) / 10.0

	envFactor := 0.0
	if node.Environment == "PROD" {
		envFactor = 1.0
	}

	blastRadius := degree / 5.0

	return 0.3*centrality +
		0.3*criticality +
		0.2*envFactor +
		0.2*blastRadius
}
