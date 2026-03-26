package costoptimizer

import "github.com/diya-suryawanshi/cloud/graph-engine/graph"

func ComputeGraphImpact(g *graph.Graph, nodeID string) float64 {
	node := g.Nodes[nodeID]

	inDegree := 0
	for _, edges := range g.Adjacency {
		for _, e := range edges {
			if e.To == nodeID {
				inDegree++
			}
		}
	}

	outDegree := len(g.Adjacency[nodeID])

	impact := float64(inDegree*2 + outDegree)

	// Environment sensitivity
	if node.Environment == "PROD" {
		impact += 2.0
	}

	// Exposure sensitivity
	if node.Exposure == "PUBLIC" {
		impact += 1.5
	}

	// Resource-type sensitivity
	switch node.Type {
	case "DATABASE":
		impact += 3.0
	case "IAM_ROLE":
		impact += 2.5
	case "LOAD_BALANCER":
		impact += 2.0
	case "COMPUTE":
		impact += 1.0
	case "OBJECT_STORAGE":
		impact += 1.0
	}

	return impact
}
