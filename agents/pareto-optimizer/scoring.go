package paretooptimizer

import (
	"math"

	"github.com/diya-suryawanshi/cloud/graph-engine/graph"
)

func ScoreAction(
	g *graph.Graph,
	action Action,
	w Weights,
	minSavings, maxSavings, minRisk, maxRisk float64,
) float64 {

	savings := -action.CostDelta
	normSavings := Normalize(savings, minSavings, maxSavings)
	normRisk := Normalize(action.RiskReduction, minRisk, maxRisk)

	penalty := ConstraintPenalty(g, action)

	score :=
		w.RiskWeight*normRisk +
			w.CostWeight*normSavings -
			w.Penalty*penalty

	return math.Round(score*100) / 100
}
