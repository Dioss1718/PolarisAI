package actiongenerator

import (
	candidategenerator "github.com/diya-suryawanshi/cloud/agents/candidate-generator"
	"github.com/diya-suryawanshi/cloud/graph-engine/graph"
)

func IsValid(g *graph.Graph, c candidategenerator.Candidate, variant string) bool {

	node := g.Nodes[c.NodeID]

	if node.Environment == "PROD" && variant == "AGGRESSIVE" {
		return false
	}

	if c.Centrality > 0.8 && c.ActionType == "TERMINATE" {
		return false
	}

	return true
}
