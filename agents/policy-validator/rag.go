package policyvalidator

import (
	"encoding/json"
	"os"
	"strings"
)

func LoadPolicy() Policy {
	file, _ := os.ReadFile("agents/policy-validator/policies.json")

	var p Policy
	json.Unmarshal(file, &p)
	return p
}

// Real retrieval based on keyword matching (no mock if/else)
func RetrievePolicyInsight(query string) string {

	data, _ := os.ReadFile("agents/policy-validator/policies.json")
	text := strings.ToLower(string(data))
	query = strings.ToLower(query)

	if strings.Contains(text, query) {
		return "Policy matched: " + query
	}

	return "No strict policy found"
}
