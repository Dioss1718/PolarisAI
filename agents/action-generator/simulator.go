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
	disruption := degree * 0.2

	switch c.ActionType {

	case "TERMINATE":
		return -c.BaseCost, c.BaseRisk * 0.3, disruption

	case "DOWNSIZE":
		if variant == "SMALL" {
			return -c.BaseCost * 0.2, 0, disruption * 0.5
		}
		if variant == "MEDIUM" {
			return -c.BaseCost * 0.4, 0, disruption * 0.7
		}
		return -c.BaseCost * 0.6, c.BaseRisk * 0.1, disruption

	case "SECURE":
		return c.BaseCost * 0.1, -c.BaseRisk * 0.6, disruption * 0.3
	}

	return 0, 0, 0
}
