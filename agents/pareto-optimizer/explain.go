package paretooptimizer

import "fmt"

func GenerateExplanation(a Action, score float64, penalty float64) string {
	return fmt.Sprintf(
		"Selected: %s on %s | RiskReduction=%.2f | CostImpact=%.2f | Penalty=%.2f | FinalScore=%.2f",
		a.ActionType,
		a.NodeID,
		a.RiskReduction,
		a.CostDelta,
		penalty,
		score,
	)
}
