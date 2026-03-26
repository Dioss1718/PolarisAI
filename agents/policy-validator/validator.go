package policyvalidator

import (
	"math"
	"strings"
)

func contains(s string, sub string) bool {
	return strings.Contains(strings.ToUpper(s), sub)
}

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

	// SLA
	if env == "PROD" && p.NoTerminateProd && contains(action.Action, "TERMINATE") {
		scores.SLA = 0.1
	} else if contains(action.Action, "DOWNSIZE") && centrality > 0.6 {
		scores.SLA = 0.5
	} else {
		scores.SLA = 1.0
	}

	// Security
	scores.Security = math.Max(0, 1-(risk/10))
	if contains(action.Action, "SECURE") || contains(action.Action, "RESTRICT") || contains(action.Action, "PATCH") {
		scores.Security += 0.2
		if scores.Security > 1 {
			scores.Security = 1
		}
	}

	// Compliance
	if p.NoPublicDB && nodeType == "DATABASE" && exposure == "PUBLIC" {
		scores.Compliance = 0.1
	} else if p.EncryptionRequired && (nodeType == "DATABASE" || nodeType == "OBJECT_STORAGE") && contains(action.Action, "TERMINATE") {
		scores.Compliance = 0.7
	} else {
		scores.Compliance = 1.0
	}

	// Blast
	scores.Blast = math.Max(0, 1-centrality)

	return scores
}
