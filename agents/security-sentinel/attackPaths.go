package securitysentinel

import (
	"github.com/diya-suryawanshi/cloud/graph-engine/graph"
)

func FindAttackPaths(g *graph.Graph) [][]string {
	var allPaths [][]string

	for _, node := range g.Nodes {

		// Entry points = PUBLIC nodes
		if node.Exposure == "PUBLIC" {
			paths := BFS(g, node.ID)
			allPaths = append(allPaths, paths...)
		}
	}

	return allPaths
}

func isSensitiveNode(g *graph.Graph, nodeID string) bool {
	node := g.Nodes[nodeID]

	// Only high-value targets
	if node.Type == "DATABASE" {
		return true
	}

	// Admin-level IAM
	if node.Type == "IAM_ROLE" {
		for _, flag := range node.Compliance {
			if flag == "ADMIN_ACCESS" {
				return true
			}
		}
	}

	return false
}

func CalculatePathRisk(g *graph.Graph, path []string) float64 {
	risk := 0.0

	for _, nodeID := range path {
		node := g.Nodes[nodeID]
		risk += float64(node.Criticality)

		if node.Exposure == "PUBLIC" {
			risk += 5
		}
	}

	return risk
}
