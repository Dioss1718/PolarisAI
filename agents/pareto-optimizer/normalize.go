package paretooptimizer

import "math"

func Normalize(value, min, max float64) float64 {
	if max-min == 0 {
		return 0
	}
	return (value - min) / (max - min)
}

func ComputeBounds(actions []Action) (float64, float64, float64, float64) {
	minCost, maxCost := math.MaxFloat64, -math.MaxFloat64
	minRisk, maxRisk := math.MaxFloat64, -math.MaxFloat64

	for _, a := range actions {
		if a.CostDelta < minCost {
			minCost = a.CostDelta
		}
		if a.CostDelta > maxCost {
			maxCost = a.CostDelta
		}
		if a.RiskReduction < minRisk {
			minRisk = a.RiskReduction
		}
		if a.RiskReduction > maxRisk {
			maxRisk = a.RiskReduction
		}
	}

	return minCost, maxCost, minRisk, maxRisk
}
