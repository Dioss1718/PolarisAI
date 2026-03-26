package paretooptimizer

import "github.com/diya-suryawanshi/cloud/graph-engine/graph"

func RunParetoOptimizer(
	g *graph.Graph,
	actions []Action,
	weights Weights,
) []Decision {
	return RunNegotiation(g, actions, weights)
}
