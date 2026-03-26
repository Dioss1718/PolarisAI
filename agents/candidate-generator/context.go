package candidategenerator

import "github.com/diya-suryawanshi/cloud/graph-engine/graph"

// ComputeNodeCentrality approximates operational importance
// using both incoming and outgoing dependencies.
func ComputeNodeCentrality(g *graph.Graph, nodeID string) float64 {
	inDegree := 0
	for _, edges := range g.Adjacency {
		for _, e := range edges {
			if e.To == nodeID {
				inDegree++
			}
		}
	}

	outDegree := len(g.Adjacency[nodeID])

	centrality := float64(inDegree*2+outDegree) / 10.0
	if centrality > 1.0 {
		return 1.0
	}
	return centrality
}
