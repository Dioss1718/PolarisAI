package securitysentinel

import (
	"github.com/diya-suryawanshi/cloud/graph-engine/graph"
)

// Approximate Betweenness Centrality (MVP safe)
func ComputeCentrality(g *graph.Graph) map[string]float64 {
	centrality := make(map[string]float64)

	for source := range g.Nodes {
		paths := BFSWithTarget(g, source)

		for _, path := range paths {
			for _, node := range path {
				if node != source {
					centrality[node] += 1
				}
			}
		}
	}

	// Normalize
	max := 1.0
	for _, v := range centrality {
		if v > max {
			max = v
		}
	}

	for k, v := range centrality {
		centrality[k] = v / max
	}

	return centrality
}
