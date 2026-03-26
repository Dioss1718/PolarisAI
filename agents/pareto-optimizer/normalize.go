package paretooptimizer

import "math"

func Normalize(value, min, max float64) float64 {
	if max-min == 0 {
		return 0
	}
	return (value - min) / (max - min)
}

// Compute bounds in the actual optimization space:
// savings = -CostDelta (higher is better)
// risk reduction = RiskReduction (higher is better)
func ComputeBounds(actions []Action) (float64, float64, float64, float64) {
	minSavings, maxSavings := math.MaxFloat64, -math.MaxFloat64
	minRisk, maxRisk := math.MaxFloat64, -math.MaxFloat64

	for _, a := range actions {
		savings := -a.CostDelta

		if savings < minSavings {
			minSavings = savings
		}
		if savings > maxSavings {
			maxSavings = savings
		}
		if a.RiskReduction < minRisk {
			minRisk = a.RiskReduction
		}
		if a.RiskReduction > maxRisk {
			maxRisk = a.RiskReduction
		}
	}

	return minSavings, maxSavings, minRisk, maxRisk
}
