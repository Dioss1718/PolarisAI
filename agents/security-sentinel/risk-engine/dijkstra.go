package riskengine

import (
	"math"

	"github.com/diya-suryawanshi/cloud/graph-engine/graph"
)

// ATTACKER PERSPECTIVE
// "How easy is it for attacker to reach this node?"
// Lower distance = easier to reach = more dangerous
func Dijkstra(g *graph.Graph, start string) map[string]float64 {
	dist := make(map[string]float64)
	visited := make(map[string]bool)

	for id := range g.Nodes {
		dist[id] = math.Inf(1)
	}
	dist[start] = 0

	for {
		minNode, minDist := "", math.Inf(1)
		for node, d := range dist {
			if !visited[node] && d < minDist {
				minDist = d
				minNode = node
			}
		}
		if minNode == "" {
			break
		}
		visited[minNode] = true

		for _, edge := range g.Adjacency[minNode] {
			weight := float64(edge.Weight)
			if g.Nodes[edge.To].Exposure == "PUBLIC" {
				weight += 5 // PUBLIC = easier target for attacker
			}
			if dist[minNode]+weight < dist[edge.To] {
				dist[edge.To] = dist[minNode] + weight
			}
		}
	}
	return dist
}

// AttackerReachability — how close is this node to any entry point
// Entry points = PUBLIC nodes + SECURITY_GROUPs
// Returns 0-10 score. Higher = easier for attacker to reach.
func AttackerReachability(g *graph.Graph, nodeID string) float64 {
	minD := math.Inf(1)

	for sid := range g.Nodes {
		s := g.Nodes[sid]
		if s.Exposure != "PUBLIC" && s.Type != "SECURITY_GROUP" {
			continue
		}
		dists := Dijkstra(g, sid)
		if d, ok := dists[nodeID]; ok && d < minD {
			minD = d
		}
	}

	if math.IsInf(minD, 1) {
		return 0.0 // unreachable from any entry = safe
	}

	// Closer to entry = higher score
	return math.Max(0, 50.0-minD) / 50.0 * 10.0
}

// CalculatePathRisk — total risk of one full attack path
// Used by security-sentinel attack_path.go
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
