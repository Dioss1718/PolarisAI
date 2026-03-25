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

	degree := float64(len(g.Adjacency[c.NodeID]))

	// 🔥 Graph-aware disruption
	dependencyImpact := degree * 0.3

	// 🔥 Environment sensitivity
	envMultiplier := 1.0
	if c.Env == "PROD" {
		envMultiplier = 1.5
	}

	disruption := dependencyImpact * envMultiplier

	switch c.ActionType {

	case "TERMINATE":
		riskReduction := c.BaseRisk * 0.3

		// 🔥 More realistic: high central nodes = more disruption
		disruption *= (1 + c.Centrality)

		return -c.BaseCost, riskReduction, disruption

	case "DOWNSIZE":
		if variant == "SMALL" {
			return -c.BaseCost * 0.2, 0, disruption * 0.4
		}
		if variant == "MEDIUM" {
			return -c.BaseCost * 0.4, 0, disruption * 0.6
		}

		// AGGRESSIVE
		return -c.BaseCost * 0.6, c.BaseRisk * 0.1, disruption

	case "SECURE":
		// 🔥 security reduces risk but may add latency (small disruption)
		return c.BaseCost * 0.1, -c.BaseRisk * 0.6, disruption * 0.3
	}

	return 0, 0, 0
}
