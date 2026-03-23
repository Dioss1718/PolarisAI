package riskengine

import (
	"math"

	securitysentinel "github.com/diya-suryawanshi/cloud/agents/security-sentinel"
	"github.com/diya-suryawanshi/cloud/graph-engine/graph"
)

const (
	Alpha = 0.4 // criticality
	Beta  = 0.3 // exposure
	Gamma = 0.2 // centrality
	Delta = 0.1 // dijkstra reachability
)

// --- DIJKSTRA ---
// Finds minimum risk distance from start node to all other nodes
func Dijkstra(g *graph.Graph, start string) map[string]float64 {
	dist := make(map[string]float64)
	visited := make(map[string]bool)

	for id := range g.Nodes {
		dist[id] = math.Inf(1)
	}
	dist[start] = 0

	for {
		minNode, minDist := "", math.Inf(1)
		for n, d := range dist {
			if !visited[n] && d < minDist {
				minDist = d
				minNode = n
			}
		}
		if minNode == "" {
			break
		}
		visited[minNode] = true

		for _, edge := range g.Adjacency[minNode] {
			t := g.Nodes[edge.To]
			w := float64(edge.Weight) + float64(t.Criticality)*0.5
			if t.Exposure == "PUBLIC" {
				w += 5.0
			}
			if dist[minNode]+w < dist[edge.To] {
				dist[edge.To] = dist[minNode] + w
			}
		}
	}
	return dist
}

// --- DIJKSTRA REACHABILITY ---
// How close is this node to any PUBLIC/SECURITY_GROUP entry point
func dijkstraContribution(nodeID string, g *graph.Graph) float64 {
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
		return 0.0
	}
	return math.Max(0, 50.0-minD) / 50.0 * 10.0
}

// --- RISK SCORE ---
// Combines centrality (from sentinel) + dijkstra + criticality + exposure
func ComputeNodeRisk(g *graph.Graph) map[string]float64 {
	centrality := securitysentinel.ComputeCentrality(g)
	scores := make(map[string]float64)

	for id, node := range g.Nodes {
		vuln := Alpha * float64(node.Criticality)

		exp := 0.0
		if node.Exposure == "PUBLIC" {
			exp = Beta * 10.0
		}

		cen := Gamma * centrality[id] * 10.0
		reach := Delta * dijkstraContribution(id, g)

		scores[id] = math.Round((vuln+exp+cen+reach)*100) / 100
	}
	return scores
}
