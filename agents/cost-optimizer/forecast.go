package costoptimizer

// Placeholder for future TFT model
func ForecastCost(current float64, utilization float64) float64 {

	// Simple heuristic growth model (MVP hook)
	if utilization < 20 {
		return current * 1.3
	}

	return current * 1.1
}
