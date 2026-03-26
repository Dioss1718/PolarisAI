package gitops

import (
	"fmt"
	"os"
)

type Plugin struct{}

func validateGitHubEnv() error {
	if os.Getenv("GITHUB_TOKEN") == "" {
		return fmt.Errorf("GITHUB_TOKEN is missing")
	}
	if os.Getenv("GITHUB_REPO") == "" {
		return fmt.Errorf("GITHUB_REPO is missing")
	}
	return nil
}

func (p *Plugin) Run(
	graph interface{},
	decisions []Decision,
	nodeRisks map[string]float64,
) ([]PRResponse, error) {

	if err := validateGitHubEnv(); err != nil {
		return nil, err
	}

	current := graph.(*Graph)

	_, prs := RunGitOps(current, decisions, nodeRisks)

	return prs, nil
}
