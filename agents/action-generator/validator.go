package actiongenerator

import (
	candidategenerator "github.com/diya-suryawanshi/cloud/agents/candidate-generator"
	"github.com/diya-suryawanshi/cloud/graph-engine/graph"
)

func IsValid(g *graph.Graph, c candidategenerator.Candidate, variant string) bool {

	node := g.Nodes[c.NodeID]

	// 🔒 Strong PROD protection
	if node.Environment == "PROD" && (variant == "AGGRESSIVE" || c.ActionType == "TERMINATE") {
		return false
	}

	// 🔒 High central nodes → no destructive ops
	if c.Centrality > 0.8 && c.ActionType == "TERMINATE" {
		return false
	}

	// 🔒 Avoid forceful actions on risky nodes
	if c.BaseRisk > 8 && variant == "FORCE" {
		return false
	}

	return true
}
