package securitysentinel

import (
	"github.com/diya-suryawanshi/cloud/graph-engine/graph"
)

func BFS(g *graph.Graph, start string) [][]string {
	var paths [][]string

	queue := [][]string{{start}}
	visited := make(map[string]bool)

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]

		lastNode := path[len(path)-1]

		if visited[lastNode] {
			continue
		}
		visited[lastNode] = true

		neighbors := g.GetNeighbors(lastNode)

		for _, edge := range neighbors {
			newPath := append([]string{}, path...)
			newPath = append(newPath, edge.To)

			if isSensitiveNode(g, edge.To) {
				paths = append(paths, newPath)
				continue
			}

			queue = append(queue, newPath)
		}
	}

	return paths
}
