package paretooptimizer

import "github.com/diya-suryawanshi/cloud/graph-engine/graph"

func ScoreAction(
	g *graph.Graph,
	action Action,
	w Weights,
	minCost, maxCost, minRisk, maxRisk float64,
) float64 {

	normCost := Normalize(-action.CostDelta, minCost, maxCost)
	normRisk := Normalize(action.RiskReduction, minRisk, maxRisk)

	penalty := ConstraintPenalty(g, action)

	score :=
		w.RiskWeight*normRisk +
			w.CostWeight*normCost -
			w.Penalty*penalty

	return score
}
