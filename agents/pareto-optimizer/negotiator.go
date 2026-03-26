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

	// Step 1: Filter valid actions
	var valid []Action
	for _, a := range actions {
		if IsValid(g, a) {
			valid = append(valid, a)
		}
	}

	if len(valid) == 0 {
		return []Decision{}
	}

	// Step 2: Pareto front
	pareto := FilterPareto(valid)

	minCost, maxCost, minRisk, maxRisk := ComputeBounds(pareto)

	var decisions []Decision

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

		decisions = append(decisions, Decision{
			NodeID: a.NodeID,
			Action: a.ActionType,
			Score:  score,
			Reason: GenerateExplanation(a, score, penalty),
		})
	}

	sort.Slice(decisions, func(i, j int) bool {
		return decisions[i].Score > decisions[j].Score
	})

	return decisions
}
