package paretooptimizer

import (
	"sort"

	"github.com/diya-suryawanshi/cloud/graph-engine/graph"
)

func RunNegotiation(
	g *graph.Graph,
	actions []Action,
) []Decision {

	// Step 1: Filter invalid actions
	var valid []Action
	for _, a := range actions {
		if IsValid(g, a) {
			valid = append(valid, a)
		}
	}

	if len(valid) == 0 {
		return []Decision{}
	}

	// Step 2: Pareto Front (per-node)
	pareto := FilterPareto(valid)
	if len(pareto) == 0 {
		return []Decision{}
	}

	// Step 3: Bounds in normalized savings/risk space
	minSavings, maxSavings, minRisk, maxRisk := ComputeBounds(pareto)

	// Step 4: Initial weights
	weights := Weights{
		RiskWeight: 0.55,
		CostWeight: 0.45,
		Penalty:    0.70,
	}

	// Step 5: Lightweight learning / adaptation loop
	for i := 0; i < 5; i++ {
		var total float64

		for _, a := range pareto {
			score := ScoreAction(g, a, weights, minSavings, maxSavings, minRisk, maxRisk)
			total += score
		}

		avg := total / float64(len(pareto))
		weights = UpdateWeights(weights, avg)
	}

	// Step 6: pick best negotiated action per node
	bestByNode := make(map[string]Decision)

	for _, a := range pareto {
		penalty := ConstraintPenalty(g, a)
		score := ScoreAction(g, a, weights, minSavings, maxSavings, minRisk, maxRisk)

		decision := Decision{
			NodeID: a.NodeID,
			Action: a.ActionType,
			Score:  score,
			Reason: GenerateExplanation(a, score, penalty),
		}

		if existing, ok := bestByNode[a.NodeID]; !ok || decision.Score > existing.Score {
			bestByNode[a.NodeID] = decision
		}
	}

	var decisions []Decision
	for _, d := range bestByNode {
		decisions = append(decisions, d)
	}

	// Step 7: rank globally
	sort.Slice(decisions, func(i, j int) bool {
		return decisions[i].Score > decisions[j].Score
	})

	return decisions
}
