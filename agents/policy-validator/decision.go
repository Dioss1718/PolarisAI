package policyvalidator

func ComputeFinalDecision(
	input InputDecision,
	scores ValidationScores,
	env string,
) ValidatedDecision {

	wSLA := 0.3
	wSec := 0.3
	wComp := 0.2
	wBlast := 0.2

	if env == "PROD" {
		wSLA += 0.1
		wBlast += 0.1
		wSec -= 0.05
		wComp -= 0.05
	}

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
