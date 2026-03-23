package costoptimizer

func GenerateCostCandidates(signals []CostSignal) []CostCandidate {

	var candidates []CostCandidate

	for _, s := range signals {

		var action string

		if s.OptimizationPotential > 0.7 && s.GraphImpact < 2 {
			action = "DELETE"
		} else if s.OptimizationPotential > 0.5 {
			action = "SCHEDULE_STOP"
		} else {
			action = "RIGHTSIZE"
		}

		delta := EstimateCostDelta(s.Utilization, s.CurrentCost)

		score := delta * s.Confidence

		candidates = append(candidates, CostCandidate{
			NodeID:     s.NodeID,
			ActionType: action,
			DeltaCost:  delta,
			Score:      score,
		})
	}

	return candidates
}
