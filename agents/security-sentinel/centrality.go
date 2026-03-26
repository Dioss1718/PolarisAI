package securitysentinel

import "github.com/diya-suryawanshi/cloud/graph-engine/graph"

// ComputeCentrality approximates betweenness centrality by
// counting how often a node appears as an internal step
// across discovered attack paths.
func ComputeCentrality(g *graph.Graph) map[string]float64 {
	centrality := make(map[string]float64)

	paths := FindAttackPaths(g)

	for _, path := range paths {
		if len(path) <= 2 {
			continue
		}

		// only count internal nodes, not source or target
		for i := 1; i < len(path)-1; i++ {
			centrality[path[i]] += 1
		}
	}

	maxVal := 1.0
	for _, v := range centrality {
		if v > maxVal {
			maxVal = v
		}
	}

	for k, v := range centrality {
		centrality[k] = v / maxVal
	}

	return centrality
}
