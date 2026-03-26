package feedback

import "math"

// Convert outcome → reward (0 to 1)
func Evaluate(costDelta, riskReduction float64) float64 {

	costImpact := -costDelta
	riskImpact := riskReduction

	score :=
		0.5*(costImpact/(1+math.Abs(costImpact))) +
			0.5*(riskImpact/(1+math.Abs(riskImpact)))

	if score < 0 {
		return 0
	}
	if score > 1 {
		return 1
	}
	return score
}
