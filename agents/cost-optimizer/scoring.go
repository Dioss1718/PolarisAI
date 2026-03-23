package costoptimizer

func computeAdaptiveWeights(cost float64, utilization float64, env string) (float64, float64, float64) {

	// dynamic weighting (NOT static)
	wCost := 0.3
	wUtil := 0.4
	wGraph := 0.3

	// High cost resources → prioritize cost
	if cost > 100 {
		wCost += 0.1
	}

	// Low utilization → prioritize utilization
	if utilization < 30 {
		wUtil += 0.1
	}

	// PROD → penalize aggressive optimization
	if env == "PROD" {
		wGraph += 0.1
	}

	return wCost, wUtil, wGraph
}

func ComputeScore(cost float64, wasteRatio float64, graphImpact float64, env string) float64 {

	wCost, wUtil, wGraph := computeAdaptiveWeights(cost, wasteRatio*100, env)

	// Normalize graph impact (inverse effect)
	graphFactor := 1 / (1 + graphImpact)

	score :=
		wCost*(cost/100) +
			wUtil*wasteRatio +
			wGraph*graphFactor

	return score
}

func ComputeConfidence(graphImpact float64, env string) float64 {

	conf := 1.0 - (graphImpact * 0.05)

	if env == "PROD" {
		conf -= 0.2
	}

	if conf < 0.3 {
		conf = 0.3
	}

	return conf
}
