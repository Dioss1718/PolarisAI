package gitops

import (
	"strings"

	modelspkg "github.com/diya-suryawanshi/cloud/graph-engine/models"
)

func GenerateProposedGraph(current *Graph, d Decision) *Graph {
	newGraph := CloneGraph(current)

	action := normalizeProposedAction(d.FinalAction)

	node, ok := newGraph.Nodes[d.NodeID]
	if !ok {
		return newGraph
	}

	switch action {
	case "TERMINATE_SAFE", "TERMINATE_FORCE":
		removeNodeAndEdges(newGraph, d.NodeID)
		return newGraph

	case "DOWNSIZE_SMALL":
		node.Utilization = clamp(node.Utilization * 0.85)
		node.Cost = clampCost(node.Cost * 0.85)

	case "DOWNSIZE_MEDIUM":
		node.Utilization = clamp(node.Utilization * 0.70)
		node.Cost = clampCost(node.Cost * 0.70)

	case "DOWNSIZE_AGGRESSIVE":
		node.Utilization = clamp(node.Utilization * 0.55)
		node.Cost = clampCost(node.Cost * 0.55)

	case "SECURE_PATCH":
		if node.Exposure == "PUBLIC" {
			node.Exposure = "INTERNAL"
		}
		node.Compliance = removeComplianceFlags(
			node.Compliance,
			"OPEN_PORT",
			"PORT_0_0_0_0",
		)
		node.Cost = clampCost(node.Cost * 1.05)

	case "SECURE_RESTRICT":
		node.Exposure = "PRIVATE"
		node.Compliance = removeComplianceFlags(
			node.Compliance,
			"PUBLIC_BUCKET",
			"OPEN_PORT",
			"PORT_0_0_0_0",
			"ADMIN_ACCESS",
			"IAM_OVERPRIVILEGED",
		)
		node.Cost = clampCost(node.Cost * 1.03)
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
		complianceCopy := append([]string{}, v.Compliance...)
		v.Compliance = complianceCopy
		newGraph.Nodes[k] = v
	}

	for k, edges := range g.Adjacency {
		cloned := make([]modelspkg.Edge, len(edges))
		copy(cloned, edges)
		newGraph.Adjacency[k] = cloned
	}

	return newGraph
}

func normalizeProposedAction(action string) string {
	return strings.TrimPrefix(strings.TrimSpace(action), "SAFE_")
}

func removeNodeAndEdges(g *Graph, nodeID string) {
	delete(g.Nodes, nodeID)
	delete(g.Adjacency, nodeID)

	filteredEdges := make([]modelspkg.Edge, 0, len(g.Edges))
	for _, edge := range g.Edges {
		if edge.From == nodeID || edge.To == nodeID {
			continue
		}
		filteredEdges = append(filteredEdges, edge)
	}
	g.Edges = filteredEdges

	newAdj := make(map[string][]modelspkg.Edge)
	for from, edges := range g.Adjacency {
		filtered := make([]modelspkg.Edge, 0, len(edges))
		for _, edge := range edges {
			if edge.From == nodeID || edge.To == nodeID {
				continue
			}
			filtered = append(filtered, edge)
		}
		newAdj[from] = filtered
	}
	g.Adjacency = newAdj
}

func removeComplianceFlags(items []string, forbidden ...string) []string {
	blocked := map[string]bool{}
	for _, f := range forbidden {
		blocked[f] = true
	}

	out := make([]string, 0, len(items))
	for _, item := range items {
		if !blocked[item] {
			out = append(out, item)
		}
	}
	return out
}

func clamp(v float64) float64 {
	if v < 0 {
		return 0
	}
	if v > 1 {
		return 1
	}
	return v
}

func clampCost(v float64) float64 {
	if v < 0 {
		return 0
	}
	return v
}
