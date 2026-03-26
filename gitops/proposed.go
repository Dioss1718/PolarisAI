package gitops

import (
	modelspkg "github.com/diya-suryawanshi/cloud/graph-engine/models"
)

func GenerateProposedGraph(current *Graph, d Decision) *Graph {
	newGraph := CloneGraph(current)

	node, ok := newGraph.Nodes[d.NodeID]
	if !ok {
		return newGraph
	}

	switch d.FinalAction {
	case "TERMINATE_SAFE", "TERMINATE_FORCE":

		node.Exposure = "PRIVATE"
		node.Utilization = 0
		node.Cost = 0

	case "DOWNSIZE_SMALL":
		node.Utilization = node.Utilization * 0.85
		node.Cost = node.Cost * 0.85

	case "DOWNSIZE_MEDIUM":
		node.Utilization = node.Utilization * 0.70
		node.Cost = node.Cost * 0.70

	case "DOWNSIZE_AGGRESSIVE":
		node.Utilization = node.Utilization * 0.55
		node.Cost = node.Cost * 0.55

	case "SECURE_PATCH":
		node.Exposure = "INTERNAL"
		node.Compliance = appendIfMissing(node.Compliance, "PATCHED")
		node.Cost = node.Cost * 1.05

	case "SECURE_RESTRICT":
		node.Exposure = "PRIVATE"
		node.Compliance = appendIfMissing(node.Compliance, "ACCESS_RESTRICTED")
		node.Cost = node.Cost * 1.03
	}

	newGraph.Nodes[d.NodeID] = node
	return newGraph
}

func GenerateFullProposedGraph(current *Graph, decisions []Decision) *Graph {
	newGraph := CloneGraph(current)

	for _, d := range decisions {
		newGraph = GenerateProposedGraph(newGraph, d)
	}

	return newGraph
}

func CloneGraph(g *Graph) *Graph {
	newGraph := &Graph{
		Nodes:     make(map[string]modelspkg.Node),
		Edges:     make([]modelspkg.Edge, len(g.Edges)),
		Adjacency: make(map[string][]modelspkg.Edge),
	}

	copy(newGraph.Edges, g.Edges)

	for k, v := range g.Nodes {
		newGraph.Nodes[k] = v
	}

	for k, edges := range g.Adjacency {
		cloned := make([]modelspkg.Edge, len(edges))
		copy(cloned, edges)
		newGraph.Adjacency[k] = cloned
	}

	return newGraph
}

func appendIfMissing(items []string, value string) []string {
	for _, v := range items {
		if v == value {
			return items
		}
	}
	return append(items, value)
}
