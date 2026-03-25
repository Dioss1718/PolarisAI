package policyvalidator

import "fmt"

func Audit(dec ValidatedDecision) {
	fmt.Printf(
		"[AUDIT] Node=%s | Action=%s | Status=%s | Score=%.2f\n",
		dec.NodeID, dec.FinalAction, dec.Status, dec.Score,
	)
}
