package actiongenerator

import "math"

func Score(costDelta, riskReduction, disruption float64) float64 {

	// 🔥 Normalize values
	costScore := -costDelta
	riskScore := riskReduction

	// 🔥 Non-linear penalty (important for top systems)
	disruptionPenalty := math.Pow(disruption, 1.2)

	// 🔥 Confidence estimation
	confidence := 1.0 / (1.0 + disruption)

	score :=
		(0.4 * riskScore) +
			(0.4 * costScore) +
			(0.2 * confidence) -
			(0.3 * disruptionPenalty)

	return score
}
