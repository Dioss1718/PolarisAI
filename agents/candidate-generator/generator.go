package candidategenerator

import (
	costoptimizer "github.com/diya-suryawanshi/cloud/agents/cost-optimizer"
	"github.com/diya-suryawanshi/cloud/graph-engine/graph"
)

func GenerateCandidates(
	g *graph.Graph,
	signals []costoptimizer.CostSignal,
	risks map[string]float64,
) []Candidate {

	var candidates []Candidate

	for _, s := range signals {

		risk := risks[s.NodeID]
		node := g.Nodes[s.NodeID]

		centrality := float64(len(g.Adjacency[s.NodeID])) / 5.0

		if s.WasteRatio > 0.7 {
			candidates = append(candidates,
				NewCandidate(s.NodeID, "TERMINATE", s.Cost, risk, centrality, node.Environment))
		}

		if s.Utilization < 30 {
			candidates = append(candidates,
				NewCandidate(s.NodeID, "DOWNSIZE", s.Cost, risk, centrality, node.Environment))
		}

		if risk > 6 {
			candidates = append(candidates,
				NewCandidate(s.NodeID, "SECURE", s.Cost, risk, centrality, node.Environment))
		}
	}

	return candidates
}
