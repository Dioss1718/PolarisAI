package paretooptimizer

// Check if action A dominates B
func dominates(a, b Action) bool {
	return (a.RiskReduction >= b.RiskReduction &&
		a.CostDelta <= b.CostDelta) &&
		(a.RiskReduction > b.RiskReduction ||
			a.CostDelta < b.CostDelta)
}

// Pareto front filtering
func FilterPareto(actions []Action) []Action {
	var pareto []Action

	for i, a := range actions {
		dominated := false

		for j, b := range actions {
			if i != j && dominates(b, a) {
				dominated = true
				break
			}
		}

		if !dominated {
			pareto = append(pareto, a)
		}
	}

	return pareto
}
