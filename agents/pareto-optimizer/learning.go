package paretooptimizer

func UpdateWeights(w Weights, reward float64) Weights {
	learningRate := 0.08

	if reward > 0.75 {
		w.RiskWeight += learningRate * 0.04
		w.CostWeight += learningRate * 0.03
		w.Penalty -= learningRate * 0.02
	} else {
		w.RiskWeight += learningRate * 0.06
		w.CostWeight -= learningRate * 0.02
		w.Penalty += learningRate * 0.03
	}

	if w.Penalty < 0.3 {
		w.Penalty = 0.3
	}
	if w.Penalty > 1.0 {
		w.Penalty = 1.0
	}

	total := w.RiskWeight + w.CostWeight
	if total == 0 {
		w.RiskWeight = 0.5
		w.CostWeight = 0.5
		return w
	}

	w.RiskWeight /= total
	w.CostWeight /= total

	return w
}
