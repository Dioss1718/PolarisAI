package policyvalidator

import "fmt"

func GenerateExplanation(
	nodeID string,
	action string,
	score float64,
	reason string,
) string {
	return fmt.Sprintf(
		"Node=%s | Action=%s | Score=%.2f | Insight=%s",
		nodeID, action, score, reason,
	)
}
