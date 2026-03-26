package paretooptimizer

// Check if action A dominates B
func dominates(a, b Action) bool {
	return (a.RiskReduction >= b.RiskReduction &&
		a.CostDelta <= b.CostDelta) &&
		(a.RiskReduction > b.RiskReduction ||
			a.CostDelta < b.CostDelta)
}

// Pareto front filtering should happen per node, not globally.
func FilterPareto(actions []Action) []Action {
	grouped := make(map[string][]Action)

	for _, a := range actions {
		grouped[a.NodeID] = append(grouped[a.NodeID], a)
	}

	var pareto []Action

	for _, nodeActions := range grouped {
		for i, a := range nodeActions {
			dominated := false

			for j, b := range nodeActions {
				if i != j && dominates(b, a) {
					dominated = true
					break
				}
			}

			if !dominated {
				pareto = append(pareto, a)
			}
		}
	}

	return pareto
}
