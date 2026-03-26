package securitysentinel

import (
	"sort"
	"strings"

	"github.com/diya-suryawanshi/cloud/graph-engine/graph"
)

// FindAttackPaths discovers possible attack paths
// from attacker entry points to sensitive targets.
func FindAttackPaths(g *graph.Graph) [][]string {
	var allPaths [][]string

	for _, node := range g.Nodes {
		if isEntryPoint(node.Type, node.Exposure) {
			paths := BFSWithTarget(g, node.ID)
			allPaths = append(allPaths, paths...)
		}
	}

	return deduplicatePaths(allPaths)
}

// BFSWithTarget explores multiple valid attack paths from a start node
// to sensitive targets, without prematurely blocking alternate paths.
func BFSWithTarget(g *graph.Graph, start string) [][]string {
	var result [][]string

	queue := [][]string{{start}}
	maxDepth := len(g.Nodes) + 2

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]

		if len(path) > maxDepth {
			continue
		}

		current := path[len(path)-1]

		if isSensitiveNode(g, current) && current != start {
			result = append(result, path)
			continue
		}

		for _, neighbor := range g.Adjacency[current] {
			if containsNode(path, neighbor.To) {
				continue
			}

			newPath := append([]string{}, path...)
			newPath = append(newPath, neighbor.To)
			queue = append(queue, newPath)
		}
	}

	return result
}

func isEntryPoint(nodeType, exposure string) bool {
	return exposure == "PUBLIC" || nodeType == "SECURITY_GROUP"
}

func isSensitiveNode(g *graph.Graph, nodeID string) bool {
	node := g.Nodes[nodeID]

	if node.Type == "DATABASE" {
		return true
	}

	if node.Type == "IAM_ROLE" {
		for _, flag := range node.Compliance {
			if flag == "ADMIN_ACCESS" {
				return true
			}
		}
	}

	return false
}

func containsNode(path []string, target string) bool {
	for _, n := range path {
		if n == target {
			return true
		}
	}
	return false
}

func deduplicatePaths(paths [][]string) [][]string {
	seen := make(map[string]bool)
	var unique [][]string

	for _, path := range paths {
		key := strings.Join(path, "->")
		if !seen[key] {
			seen[key] = true
			unique = append(unique, path)
		}
	}

	sort.Slice(unique, func(i, j int) bool {
		return len(unique[i]) < len(unique[j])
	})

	return unique
}
