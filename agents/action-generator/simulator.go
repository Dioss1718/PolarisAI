package actiongenerator

import (
	candidategenerator "github.com/diya-suryawanshi/cloud/agents/candidate-generator"
	"github.com/diya-suryawanshi/cloud/graph-engine/graph"
)

func Simulate(
	g *graph.Graph,
	c candidategenerator.Candidate,
	variant string,
) (float64, float64, float64) {

	// Compute richer dependency impact: in-degree + out-degree
	inDegree := 0
	for _, edges := range g.Adjacency {
		for _, e := range edges {
			if e.To == c.NodeID {
				inDegree++
			}
		}
	}

	outDegree := len(g.Adjacency[c.NodeID])
	dependencyImpact := float64(inDegree*2+outDegree) * 0.20

	envMultiplier := 1.0
	if c.Env == "PROD" {
		envMultiplier = 1.5
	}

	disruption := dependencyImpact * envMultiplier

	switch c.ActionType {

	case "TERMINATE":
		// SAFE vs FORCE must behave differently
		if variant == "SAFE" {
			return -c.BaseCost * 0.85, c.BaseRisk * 0.20, disruption * (1.0 + c.Centrality*0.8)
		}
		// FORCE
		return -c.BaseCost, c.BaseRisk * 0.30, disruption * (1.2 + c.Centrality)

	case "DOWNSIZE":
		if variant == "SMALL" {
			return -c.BaseCost * 0.20, c.BaseRisk * 0.02, disruption * 0.35
		}
		if variant == "MEDIUM" {
			return -c.BaseCost * 0.40, c.BaseRisk * 0.05, disruption * 0.55
		}
		// AGGRESSIVE
		return -c.BaseCost * 0.60, c.BaseRisk * 0.10, disruption * 0.90

	case "SECURE":
		// PATCH: more risk reduction, slightly more cost
		if variant == "PATCH" {
			return c.BaseCost * 0.10, c.BaseRisk * 0.55, disruption * 0.25
		}
		// RESTRICT: slightly lower cost, strong security control
		if variant == "RESTRICT" {
			return c.BaseCost * 0.06, c.BaseRisk * 0.45, disruption * 0.20
		}
	}

	return 0, 0, 0
}
