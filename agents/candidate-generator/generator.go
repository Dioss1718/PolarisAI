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

		// 🔥 INTELLIGENT DECISION ENGINE

		// 1. HIGH WASTE + LOW IMPORTANCE → TERMINATE
		if s.WasteRatio > 0.75 && centrality < 0.5 {
			candidates = append(candidates,
				NewCandidate(s.NodeID, "TERMINATE", s.Cost, risk, centrality, node.Environment))
		}

		// 2. MODERATE WASTE → DOWNSIZE
		if s.Utilization < 40 && s.WasteRatio > 0.4 {
			candidates = append(candidates,
				NewCandidate(s.NodeID, "DOWNSIZE", s.Cost, risk, centrality, node.Environment))
		}

		// 3. HIGH RISK → SECURE
		if risk > 6 {
			candidates = append(candidates,
				NewCandidate(s.NodeID, "SECURE", s.Cost, risk, centrality, node.Environment))
		}

		// 4. CRITICAL NODE → AVOID TERMINATE → FORCE SAFE ACTION
		if centrality > 0.7 && risk > 5 {
			candidates = append(candidates,
				NewCandidate(s.NodeID, "SECURE", s.Cost, risk, centrality, node.Environment))
		}

		// 5. LOW COST + HIGH RISK → SECURITY PRIORITY
		if s.Cost < 50 && risk > 7 {
			candidates = append(candidates,
				NewCandidate(s.NodeID, "SECURE", s.Cost, risk, centrality, node.Environment))
		}
	}

	return candidates
}
