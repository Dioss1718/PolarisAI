package costoptimizer

import "github.com/diya-suryawanshi/cloud/graph-engine/graph"

func ComputeGraphImpact(g *graph.Graph, nodeID string) float64 {
	in, out := 0, 0

	for _, edges := range g.Adjacency {
		for _, e := range edges {
			if e.To == nodeID {
				in++
			}
		}
	}

	out = len(g.Adjacency[nodeID])

	return float64(in*2 + out)
}
