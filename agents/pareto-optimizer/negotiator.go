package paretooptimizer

import (
	"sort"

	"github.com/diya-suryawanshi/cloud/graph-engine/graph"
)

func RunNegotiation(
	g *graph.Graph,
	actions []Action,
	weights Weights,
) []Decision {

	var valid []Action
	for _, a := range actions {
		if IsValid(g, a) {
			valid = append(valid, a)
		}
	}

	if len(valid) == 0 {
		return []Decision{}
	}

	pareto := FilterPareto(valid)
	if len(pareto) == 0 {
		return []Decision{}
	}

	minCost, maxCost, minRisk, maxRisk := ComputeBounds(pareto)

	bestByNode := make(map[string]Decision)

	for _, a := range pareto {
		penalty := ConstraintPenalty(g, a)

		score := ScoreAction(
			g,
			a,
			weights,
			minCost,
			maxCost,
			minRisk,
			maxRisk,
		)

		decision := Decision{
			NodeID: a.NodeID,
			Action: a.ActionType,
			Score:  score,
			Reason: GenerateExplanation(a, score, penalty),
		}

		existing, ok := bestByNode[a.NodeID]
		if !ok || decision.Score > existing.Score {
			bestByNode[a.NodeID] = decision
		}
	}

	var decisions []Decision
	for _, d := range bestByNode {
		decisions = append(decisions, d)
	}

	sort.Slice(decisions, func(i, j int) bool {
		return decisions[i].Score > decisions[j].Score
	})

	return decisions
}
