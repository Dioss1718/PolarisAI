package costoptimizer

import "github.com/diya-suryawanshi/cloud/graph-engine/graph"

func Run(g *graph.Graph) []CostSignal {
	return Analyze(g)
}
