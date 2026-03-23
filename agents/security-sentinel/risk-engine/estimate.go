package riskengine

import (
	"math"

	securitysentinel "github.com/diya-suryawanshi/cloud/agents/security-sentinel"
	"github.com/diya-suryawanshi/cloud/graph-engine/graph"
)

const (
	Alpha = 0.4 // criticality weight
	Beta  = 0.3 // exposure weight
	Gamma = 0.2 // centrality weight
	Delta = 0.1 // attacker reachability weight
)

// OUR PERSPECTIVE
// "How many edges touch this node + their weights"
// Simple count of connections — our graph view
func EstimatePathRiskContribution(g *graph.Graph, nodeID string) float64 {
	score := 0.0
	for _, edge := range g.Edges {
		if edge.From == nodeID || edge.To == nodeID {
			score += float64(edge.Weight)
		}
	}
	return score
}

// ComputeDegreeCentrality — how connected each node is
// More connections = more central = more dangerous if compromised
func ComputeDegreeCentrality(g *graph.Graph) map[string]float64 {
	centrality := make(map[string]float64)
	for _, edge := range g.Edges {
		centrality[edge.From]++
		centrality[edge.To]++
	}
	maxVal := 1.0
	for _, v := range centrality {
		if v > maxVal {
			maxVal = v
		}
	}
	for k, v := range centrality {
		centrality[k] = v / maxVal
	}
	return centrality
}

// ComputeNodeRisk — MAIN FUNCTION
// Combines BOTH perspectives into one final score per node
//
// OUR perspective:
//   Alpha × criticality       (how important is this node)
//   Beta  × exposure          (is it public facing)
//   Gamma × centrality        (how connected in graph)
//
// ATTACKER perspective (from dijkstra.go):
//   Delta × AttackerReachability  (how easy to reach from internet)
//
// Output: map[nodeID]float64  e.g. {"aws_vm1": 7.82, "sg1": 9.10}
// This output goes directly to pareto-optimizer
func ComputeNodeRisk(g *graph.Graph) map[string]float64 {
	riskScores := make(map[string]float64)

	// Centrality from security-sentinel (betweenness centrality)
	centrality := securitysentinel.ComputeCentrality(g)

	for id, node := range g.Nodes {

		// OUR perspective components
		vulnerability := float64(node.Criticality)

		exposure := 0.0
		if node.Exposure == "PUBLIC" {
			exposure = 10
		}

		centralityScore := centrality[id] * 10

		// ATTACKER perspective component (from dijkstra.go)
		attackerReach := AttackerReachability(g, id)

		// Final combined score
		risk := Alpha*vulnerability +
			Beta*exposure +
			Gamma*centralityScore +
			Delta*attackerReach

		riskScores[id] = math.Round(risk*100) / 100
	}

	return riskScores
}
