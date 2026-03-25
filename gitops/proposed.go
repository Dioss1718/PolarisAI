package gitops

import (
	models "github.com/diya-suryawanshi/cloud/graph-engine/models"
)



func GenerateProposedGraph(current *Graph, d Decision) *Graph {

	// 🔹 Create new graph
	newGraph := &Graph{
		Nodes:     make(map[string]models.Node),
		Edges:     make([]models.Edge, len(current.Edges)),
		Adjacency: make(map[string][]models.Edge),
	}

	// 🔹 Copy edges
	copy(newGraph.Edges, current.Edges)

	// 🔹 Copy nodes
	for id, n := range current.Nodes {
		newGraph.Nodes[id] = n
	}

	// 🔹 Copy adjacency
	for k, v := range current.Adjacency {
		newSlice := make([]models.Edge, len(v))
		copy(newSlice, v)
		newGraph.Adjacency[k] = newSlice
	}

	// 🔥 APPLY DECISION
	node, exists := newGraph.Nodes[d.NodeID]
	if !exists {
		return newGraph
	}

	switch d.FinalAction {

	// 🔐 SECURITY HARDENING
	case "RESTRICT_ACCESS":
		node.Exposure = "PRIVATE"

	// 🔥 ISOLATION (simulate terminate safely)
	case "SAFE_TERMINATE":
		node.Type = "ISOLATED"

	// 🔒 ENCRYPTION
	case "ENCRYPT":
		// mark via compliance flag
		node.Compliance = append(node.Compliance, "ENCRYPTED")

	// 💰 COST OPTIMIZATION (example dynamic)
	case "RIGHTSIZE":
		node.Utilization = node.Utilization * 0.8

	default:
		// Generic transformation
		node.Type = d.FinalAction
	}

	newGraph.Nodes[d.NodeID] = node

	return newGraph
}