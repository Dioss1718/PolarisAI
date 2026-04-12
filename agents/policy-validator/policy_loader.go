package policyvalidator

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func defaultPolicy() Policy {
	return Policy{
		MaxDowntime:        0.02,
		NoTerminateProd:    true,
		NoPublicDB:         true,
		EncryptionRequired: true,
	}
}

func LoadPolicy() Policy {
	candidates := []string{
		filepath.Join("agents", "policy-validator", "policies.json"),
		filepath.Join(".", "agents", "policy-validator", "policies.json"),
		filepath.Join("..", "agents", "policy-validator", "policies.json"),
		filepath.Join("..", "..", "agents", "policy-validator", "policies.json"),
	}

	if cwd, err := os.Getwd(); err == nil {
		candidates = append(candidates,
			filepath.Join(cwd, "agents", "policy-validator", "policies.json"),
			filepath.Join(cwd, "..", "agents", "policy-validator", "policies.json"),
			filepath.Join(cwd, "..", "..", "agents", "policy-validator", "policies.json"),
		)
	}

	for _, path := range candidates {
		file, err := os.ReadFile(path)
		if err != nil {
			continue
		}

		var p Policy
		if err := json.Unmarshal(file, &p); err != nil {
			continue
		}

		return p
	}

	return defaultPolicy()
}
