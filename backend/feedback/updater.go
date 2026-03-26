package feedback

import pareto "github.com/diya-suryawanshi/cloud/agents/pareto-optimizer"

func UpdateWeights(w pareto.Weights, s Summary) pareto.Weights {

	if s.AvgReward > 0.7 {
		w.RiskWeight += 0.05
		w.CostWeight -= 0.02
		w.Penalty -= 0.02
	} else {
		w.CostWeight += 0.05
		w.RiskWeight -= 0.02
		w.Penalty += 0.02
	}

	total := w.RiskWeight + w.CostWeight
	w.RiskWeight /= total
	w.CostWeight /= total

	return w
}
