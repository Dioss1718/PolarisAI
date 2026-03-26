package costoptimizer

import "math"

func computeAdaptiveWeights(cost float64, utilization float64, env string, nodeType string, exposure string) (float64, float64, float64) {
	wCost := 0.35
	wUtil := 0.40
	wGraph := 0.25

	if cost > 100 {
		wCost += 0.10
	}

	if utilization < 30 {
		wUtil += 0.10
	}

	if env == "PROD" {
		wGraph += 0.10
		wCost -= 0.05
	}

	if nodeType == "DATABASE" || nodeType == "IAM_ROLE" {
		wGraph += 0.10
		wUtil -= 0.05
	}

	if exposure == "PUBLIC" {
		wGraph += 0.05
	}

	return normalizeWeights(wCost, wUtil, wGraph)
}

func normalizeWeights(a, b, c float64) (float64, float64, float64) {
	total := a + b + c
	if total == 0 {
		return 0.33, 0.33, 0.34
	}
	return a / total, b / total, c / total
}

func ComputeScore(cost float64, wasteRatio float64, graphImpact float64, env string, nodeType string, exposure string) float64 {
	wCost, wUtil, wGraph := computeAdaptiveWeights(cost, wasteRatio*100, env, nodeType, exposure)

	normalizedCost := math.Min(cost/250.0, 1.0)
	normalizedWaste := math.Min(math.Max(wasteRatio, 0.0), 1.0)
	graphPenalty := 1.0 / (1.0 + graphImpact)

	// Stronger non-linear shaping so very wasteful expensive nodes stand out
	costComponent := math.Sqrt(normalizedCost)
	wasteComponent := math.Pow(normalizedWaste, 1.2)

	score :=
		wCost*costComponent +
			wUtil*wasteComponent +
			wGraph*graphPenalty

	return math.Round(score*100) / 100
}

func ComputeConfidence(graphImpact float64, env string, nodeType string, wasteRatio float64) float64 {
	conf := 0.95

	// Higher graph impact lowers confidence
	conf -= graphImpact * 0.03

	// Production environments require more caution
	if env == "PROD" {
		conf -= 0.15
	}

	// Sensitive node types reduce confidence
	if nodeType == "DATABASE" || nodeType == "IAM_ROLE" {
		conf -= 0.10
	}

	// If waste is clearly high, confidence in inefficiency detection goes up slightly
	if wasteRatio > 0.75 {
		conf += 0.05
	}

	if conf < 0.25 {
		conf = 0.25
	}
	if conf > 0.99 {
		conf = 0.99
	}

	return math.Round(conf*100) / 100
}

func BuildReason(nodeType, env, exposure string, cost, utilization, graphImpact float64) string {
	reason := "Cost inefficiency detected"

	if utilization < 20 {
		reason = "Severe underutilization with high waste potential"
	} else if utilization < 40 {
		reason = "Moderate underutilization with optimization opportunity"
	}

	if cost > 100 {
		reason += "; high absolute spend"
	}

	if env == "PROD" {
		reason += "; production sensitivity considered"
	}

	if exposure == "PUBLIC" {
		reason += "; public exposure increases remediation caution"
	}

	if nodeType == "DATABASE" || nodeType == "IAM_ROLE" {
		reason += "; critical resource type"
	}

	if graphImpact > 5 {
		reason += "; downstream dependency impact is high"
	}

	return reason
}
