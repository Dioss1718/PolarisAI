package actiongenerator

import (
	candidategenerator "github.com/diya-suryawanshi/cloud/agents/candidate-generator"
	"github.com/diya-suryawanshi/cloud/graph-engine/graph"
)

func IsValid(g *graph.Graph, c candidategenerator.Candidate, variant string) bool {
	node := g.Nodes[c.NodeID]

	if node.Environment == "PROD" && (variant == "AGGRESSIVE" || c.ActionType == "TERMINATE") {
		return false
	}

	// Highly central nodes: do not allow destructive termination
	if c.Centrality > 0.8 && c.ActionType == "TERMINATE" {
		return false
	}

	// Extremely risky nodes should not use forceful destructive variants
	if c.BaseRisk > 8 && variant == "FORCE" {
		return false
	}

	// Databases and IAM roles should never be aggressively downsized
	if (node.Type == "DATABASE" || node.Type == "IAM_ROLE") && variant == "AGGRESSIVE" {
		return false
	}

	return true
}
