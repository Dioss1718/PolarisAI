package paretooptimizer

func UpdateWeights(w Weights, reward float64) Weights {

	learningRate := 0.1

	if reward > 0.7 {
		w.RiskWeight += learningRate * 0.05
		w.CostWeight += learningRate * 0.05
	} else {
		w.RiskWeight += learningRate * 0.1
		w.CostWeight -= learningRate * 0.05
	}

	// Normalize
	total := w.RiskWeight + w.CostWeight
	w.RiskWeight /= total
	w.CostWeight /= total

	return w
}
