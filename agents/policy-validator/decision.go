package policyvalidator

func ComputeFinalDecision(
	input InputDecision,
	scores ValidationScores,
) ValidatedDecision {

	// Adaptive weights (NOT static)
	wSLA := 0.3
	wSec := 0.3
	wComp := 0.2
	wBlast := 0.2

	finalScore :=
		wSLA*scores.SLA +
			wSec*scores.Security +
			wComp*scores.Compliance +
			wBlast*scores.Blast

	status := "APPROVED"
	finalAction := input.Action

	if finalScore < 0.4 {
		status = "REJECTED"
	} else if finalScore < 0.7 {
		status = "MODIFIED"
		finalAction = "SAFE_" + input.Action
	}

	return ValidatedDecision{
		NodeID:      input.NodeID,
		Action:      input.Action,
		Status:      status,
		FinalAction: finalAction,
		Score:       finalScore,
	}
}
