package candidategenerator

import (
	"fmt"

	costoptimizer "github.com/diya-suryawanshi/cloud/agents/cost-optimizer"
	"github.com/diya-suryawanshi/cloud/graph-engine/graph"
)

func GenerateCandidates(
	g *graph.Graph,
	signals []costoptimizer.CostSignal,
	risks map[string]float64,
) []Candidate {
	var candidates []Candidate
	seen := make(map[string]bool)

	for _, s := range signals {
		risk := risks[s.NodeID]
		node := g.Nodes[s.NodeID]

		centrality := ComputeNodeCentrality(g, s.NodeID)

		tryAdd := func(action string) {
			key := fmt.Sprintf("%s|%s", s.NodeID, action)
			if seen[key] {
				return
			}
			seen[key] = true
			candidates = append(candidates,
				NewCandidate(
					s.NodeID,
					action,
					s.Cost,
					risk,
					centrality,
					node.Environment,
					node.Type,
					node.Exposure,
					s.WasteRatio,
				),
			)
		}

		// 1. High waste + low structural importance → TERMINATE
		if s.WasteRatio > 0.75 &&
			centrality < 0.45 &&
			node.Environment != "PROD" &&
			node.Type != "DATABASE" &&
			node.Type != "IAM_ROLE" {
			tryAdd("TERMINATE")
		}

		// 2. Moderate waste → DOWNSIZE
		if s.Utilization < 40 && s.WasteRatio > 0.40 {
			tryAdd("DOWNSIZE")
		}

		// 3. High risk → SECURE
		if risk > 6.0 {
			tryAdd("SECURE")
		}

		// 4. Public exposure + risk → SECURE
		if node.Exposure == "PUBLIC" && risk > 5.5 {
			tryAdd("SECURE")
		}

		// 5. High centrality → prefer secure action, avoid destructive action
		if centrality > 0.65 {
			tryAdd("SECURE")
		}

		// 6. Low-cost but risky nodes are still security-relevant
		if s.Cost < 50 && risk > 7.0 {
			tryAdd("SECURE")
		}

		// 7. Expensive PROD compute → downsizing candidate
		if node.Environment == "PROD" &&
			node.Type == "COMPUTE" &&
			s.Cost > 100 &&
			s.Utilization < 50 {
			tryAdd("DOWNSIZE")
		}
	}

	return candidates
}
