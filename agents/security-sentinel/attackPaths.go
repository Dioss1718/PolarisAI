package securitysentinel

import (
	"github.com/diya-suryawanshi/cloud/graph-engine/graph"
)

// FindAttackPaths discovers all possible attack paths
// from PUBLIC entry points to sensitive nodes
func FindAttackPaths(g *graph.Graph) [][]string {
	var allPaths [][]string

	for _, node := range g.Nodes {

		// Step 1: Entry points = PUBLIC nodes
		if node.Exposure == "PUBLIC" {

			// Step 2: BFS traversal with stopping condition
			paths := BFSWithTarget(g, node.ID)

			allPaths = append(allPaths, paths...)
		}
	}

	return allPaths
}

// BFSWithTarget performs BFS and stops when reaching sensitive nodes
func BFSWithTarget(g *graph.Graph, start string) [][]string {
	var result [][]string

	queue := [][]string{{start}}
	visited := make(map[string]bool)

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]

		current := path[len(path)-1]

		if visited[current] {
			continue
		}
		visited[current] = true

		// Step 3: Stop if sensitive node reached
		if isSensitiveNode(g, current) && current != start {
			result = append(result, path)
			continue
		}

		// Traverse neighbors
		for _, neighbor := range g.Adjacency[current] {
			newPath := append([]string{}, path...)
			newPath = append(newPath, neighbor.To)
			queue = append(queue, newPath)
		}
	}

	return result
}

// isSensitiveNode identifies high-value targets
func isSensitiveNode(g *graph.Graph, nodeID string) bool {
	node := g.Nodes[nodeID]

	// High-value database
	if node.Type == "DATABASE" {
		return true
	}

	// Admin IAM role
	if node.Type == "IAM_ROLE" {
		for _, flag := range node.Compliance {
			if flag == "ADMIN_ACCESS" {
				return true
			}
		}
	}

	return false
}
