package gitops

import (
	modelspkg "github.com/diya-suryawanshi/cloud/graph-engine/models"
)

// This function applies a single decision on the current graph
func GenerateProposedGraph(current *Graph, d Decision) *Graph {

	// First we clone the original graph so that we do not modify it directly
	newGraph := CloneGraph(current)

	// Get the node on which decision needs to be applied
	node, ok := newGraph.Nodes[d.NodeID]

	// If node does not exist, simply return the graph as it is
	if !ok {
		return newGraph
	}

	// Apply action based on decision type
	switch d.FinalAction {

	// If action is terminate, we remove exposure completely
	case "TERMINATE_SAFE", "TERMINATE_FORCE":
		node.Exposure = "none"

	// If action is downsize, we reduce utilization
	case "DOWNSIZE_SMALL", "DOWNSIZE_MEDIUM":
		node.Utilization = node.Utilization * 0.7

	// If action is secure, we reduce exposure and mark compliance
	case "SECURE_PATCH", "SECURE_RESTRICT":
		node.Exposure = "low"
		node.Compliance = append(node.Compliance, "secured")
	}

	// Update the modified node back into graph
	newGraph.Nodes[d.NodeID] = node

	return newGraph
}

// This function applies multiple decisions one by one on the graph
func GenerateFullProposedGraph(current *Graph, decisions []Decision) *Graph {

	// Clone original graph
	newGraph := CloneGraph(current)

	// Apply each decision sequentially
	for _, d := range decisions {
		newGraph = GenerateProposedGraph(newGraph, d)
	}

	return newGraph
}

// This function creates a copy of the graph
func CloneGraph(g *Graph) *Graph {

	// Creating new graph structure
	newGraph := &Graph{
		Nodes:     make(map[string]modelspkg.Node),
		Edges:     g.Edges,
		Adjacency: g.Adjacency,
	}

	// Copy all nodes from original graph to new graph
	for k, v := range g.Nodes {
		newGraph.Nodes[k] = v
	}

	return newGraph
}
