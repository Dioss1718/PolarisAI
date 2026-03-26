package policyvalidator

import (
	"encoding/json"
	"os"
)

func LoadPolicy() Policy {
	file, err := os.ReadFile("agents/policy-validator/policies.json")
	if err != nil {
		return Policy{
			MaxDowntime:        0.02,
			NoTerminateProd:    true,
			NoPublicDB:         true,
			EncryptionRequired: true,
		}
	}

	var p Policy
	if err := json.Unmarshal(file, &p); err != nil {
		return Policy{
			MaxDowntime:        0.02,
			NoTerminateProd:    true,
			NoPublicDB:         true,
			EncryptionRequired: true,
		}
	}

	return p
}
