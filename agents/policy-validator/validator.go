package policyvalidator

import "math"

func ValidateAll(
	action InputDecision,
	p Policy,
	env string,
	nodeType string,
	exposure string,
	centrality float64,
	risk float64,
) ValidationScores {

	scores := ValidationScores{}

	// 🔹 SLA SCORE
	if env == "PROD" && p.NoTerminateProd && action.Action == "TERMINATE" {
		scores.SLA = 0.2
	} else {
		scores.SLA = 1.0
	}

	// 🔹 SECURITY SCORE
	scores.Security = math.Max(0, 1-(risk/10))

	// 🔹 COMPLIANCE SCORE
	if p.NoPublicDB && nodeType == "DATABASE" && exposure == "PUBLIC" {
		scores.Compliance = 0.1
	} else {
		scores.Compliance = 1.0
	}

	// 🔹 BLAST RADIUS SCORE
	scores.Blast = math.Max(0, 1-centrality)

	return scores
}
