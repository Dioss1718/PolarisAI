package riskengine

import (
	"math"

	securitysentinel "github.com/diya-suryawanshi/cloud/agents/security-sentinel"
	"github.com/diya-suryawanshi/cloud/graph-engine/graph"
)

const (
	Alpha = 0.35 // criticality
	Beta  = 0.25 // exposure
	Gamma = 0.20 // centrality
	Delta = 0.20 // attacker reachability
)

// EstimatePathRiskContribution measures how exposed a node is
// in graph structure using incident edge weights.
func EstimatePathRiskContribution(g *graph.Graph, nodeID string) float64 {
	score := 0.0
	for _, edge := range g.Edges {
		if edge.From == nodeID || edge.To == nodeID {
			score += float64(edge.Weight)
		}
	}
	return score
}

// ComputeDegreeCentrality keeps a lightweight graph connectivity view.
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

// ComputeNodeRisk combines structural and attacker-perspective risk signals.
func ComputeNodeRisk(g *graph.Graph) map[string]float64 {
	riskScores := make(map[string]float64)

	centrality := securitysentinel.ComputeCentrality(g)

	for id, node := range g.Nodes {
		criticality := float64(node.Criticality)

		exposure := 0.0
		switch node.Exposure {
		case "PUBLIC":
			exposure = 10
		case "INTERNAL":
			exposure = 4
		case "PRIVATE":
			exposure = 1
		}

		centralityScore := centrality[id] * 10
		attackerReach := AttackerReachability(g, id)

		risk := Alpha*criticality +
			Beta*exposure +
			Gamma*centralityScore +
			Delta*attackerReach

		riskScores[id] = math.Round(risk*100) / 100
	}

	return riskScores
}
