package actiongenerator

import "math"

func Score(costDelta, riskReduction, disruption float64) float64 {
	// Positive savings are better
	costScore := -costDelta
	riskScore := riskReduction

	// Non-linear disruption penalty
	disruptionPenalty := math.Pow(disruption, 1.15)

	// Confidence proxy: lower disruption → higher confidence
	confidence := 1.0 / (1.0 + disruption)

	score :=
		(0.40 * riskScore) +
			(0.35 * costScore) +
			(0.20 * confidence) -
			(0.25 * disruptionPenalty)

	return math.Round(score*100) / 100
}
