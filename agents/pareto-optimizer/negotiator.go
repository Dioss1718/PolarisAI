package paretooptimizer

import (
	"sort"

	"github.com/diya-suryawanshi/cloud/graph-engine/graph"
)

func RunNegotiation(
	g *graph.Graph,
	actions []Action,
) []Decision {

	// ✅ Step 1: Filter invalid actions
	var valid []Action
	for _, a := range actions {
		if IsValid(g, a) {
			valid = append(valid, a)
		}
	}

	if len(valid) == 0 {
		return []Decision{}
	}

	// ✅ Step 2: Pareto Front
	pareto := FilterPareto(valid)

	// ✅ Step 3: Bounds
	minCost, maxCost, minRisk, maxRisk := ComputeBounds(pareto)

	// ✅ Step 4: Initial weights
	weights := Weights{
		RiskWeight: 0.5,
		CostWeight: 0.5,
		Penalty:    0.7,
	}

	// ✅ Step 5: Learning loop
	for i := 0; i < 5; i++ {

		var total float64

		for _, a := range pareto {
			score := ScoreAction(g, a, weights, minCost, maxCost, minRisk, maxRisk)
			total += score
		}

		avg := total / float64(len(pareto))
		weights = UpdateWeights(weights, avg)
	}

	// ✅ Step 6: Final decisions
	var decisions []Decision

	for _, a := range pareto {

		penalty := ConstraintPenalty(g, a)

		score := ScoreAction(g, a, weights, minCost, maxCost, minRisk, maxRisk)

		decisions = append(decisions, Decision{
			NodeID: a.NodeID,
			Action: a.ActionType,
			Score:  score,
			Reason: GenerateExplanation(a, score, penalty),
		})
	}

	// ✅ Step 7: Rank
	sort.Slice(decisions, func(i, j int) bool {
		return decisions[i].Score > decisions[j].Score
	})

	return decisions
}
