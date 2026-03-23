package costoptimizer

import "github.com/diya-suryawanshi/cloud/graph-engine/graph"

// Outgoing dependency count
func OutDegree(g *graph.Graph, nodeID string) int {
	return len(g.Adjacency[nodeID])
}

// Incoming dependency count
func InDegree(g *graph.Graph, nodeID string) int {
	count := 0
	for _, edges := range g.Adjacency {
		for _, e := range edges {
			if e.To == nodeID {
				count++
			}
		}
	}
	return count
}

// Graph impact score (higher = more critical)
func ComputeGraphImpact(g *graph.Graph, nodeID string) float64 {
	in := InDegree(g, nodeID)
	out := OutDegree(g, nodeID)

	return float64(in*2 + out)
}
