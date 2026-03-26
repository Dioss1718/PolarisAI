package paretooptimizer

import (
	"strings"

	"github.com/diya-suryawanshi/cloud/graph-engine/graph"
)

// Hard constraints (invalid actions)
func IsValid(g *graph.Graph, action Action) bool {
	node := g.Nodes[action.NodeID]
	actionType := strings.ToUpper(action.ActionType)

	// Do not allow destructive termination in PROD
	if node.Environment == "PROD" && strings.Contains(actionType, "TERMINATE") {
		return false
	}

	// Avoid forceful destructive actions on highly critical nodes
	if node.Criticality >= 9 && strings.Contains(actionType, "FORCE") {
		return false
	}

	return true
}

// Soft penalty (trade-off cost)
func ConstraintPenalty(g *graph.Graph, action Action) float64 {
	node := g.Nodes[action.NodeID]

	inDegree := 0
	for _, edges := range g.Adjacency {
		for _, e := range edges {
			if e.To == action.NodeID {
				inDegree++
			}
		}
	}

	outDegree := len(g.Adjacency[action.NodeID])

	centrality := float64(inDegree*2+outDegree) / 10.0
	if centrality > 1.0 {
		centrality = 1.0
	}

	criticality := float64(node.Criticality) / 10.0

	envFactor := 0.0
	if node.Environment == "PROD" {
		envFactor = 1.0
	}

	exposureFactor := 0.0
	if node.Exposure == "PUBLIC" {
		exposureFactor = 0.7
	}

	destructiveFactor := 0.0
	actionType := strings.ToUpper(action.ActionType)
	if strings.Contains(actionType, "TERMINATE") {
		destructiveFactor = 0.8
	} else if strings.Contains(actionType, "AGGRESSIVE") {
		destructiveFactor = 0.5
	}

	penalty :=
		0.30*centrality +
			0.25*criticality +
			0.20*envFactor +
			0.10*exposureFactor +
			0.15*destructiveFactor

	return penalty
}
