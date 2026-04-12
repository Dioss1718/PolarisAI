package gitops

import (
	"fmt"
	"strings"

	modelspkg "github.com/diya-suryawanshi/cloud/graph-engine/models"
)

var allowedFinalActions = map[string]bool{
	"DOWNSIZE_SMALL":      true,
	"DOWNSIZE_MEDIUM":     true,
	"DOWNSIZE_AGGRESSIVE": true,
	"SECURE_PATCH":        true,
	"SECURE_RESTRICT":     true,
	"TERMINATE_SAFE":      true,
	"TERMINATE_FORCE":     true,
}

var forbiddenIaCSnippets = []string{
	"local-exec",
	"remote-exec",
	"provisioner",
	"terraform destroy",
	"user_data",
	"curl ",
	"wget ",
	"rm -rf",
	"shutdown",
	"mkfs",
	"nc ",
	"netcat",
	"bash -c",
	"powershell",
}

func ValidateDecisionForGitOps(d Decision) error {
	action := normalizeActionName(d.FinalAction)
	if !allowedFinalActions[action] {
		return fmt.Errorf("gitops blocked: final action %s is not allowlisted", d.FinalAction)
	}
	return nil
}

func ValidateDiffForGitOps(diff Diff) error {
	if diff.NodeID == "" {
		return fmt.Errorf("gitops blocked: missing diff node id")
	}
	if len(diff.ChangeSet) == 0 {
		return fmt.Errorf("gitops blocked: empty change set")
	}
	return nil
}

func ValidateInfraCodeForGitOps(code InfraCode) error {
	if strings.TrimSpace(code.Content) == "" {
		return fmt.Errorf("gitops blocked: empty IaC output")
	}

	format := strings.ToLower(strings.TrimSpace(code.Format))
	if format != "" && format != "terraform" {
		return fmt.Errorf("gitops blocked: only terraform output is allowed in MVP")
	}

	lower := strings.ToLower(code.Content)
	for _, bad := range forbiddenIaCSnippets {
		if strings.Contains(lower, bad) {
			return fmt.Errorf("gitops blocked: forbidden IaC pattern detected: %s", bad)
		}
	}

	return nil
}

func ValidateNodePolicyForGitOps(diff Diff, d Decision) error {
	rawNode, ok := diff.OldState["node"]
	if !ok {
		return fmt.Errorf("gitops blocked: old node state missing")
	}

	node, ok := rawNode.(modelspkg.Node)
	if !ok {
		return fmt.Errorf("gitops blocked: invalid old node state type")
	}

	action := normalizeActionName(d.FinalAction)

	if node.Environment == "prod" && action == "TERMINATE_FORCE" {
		return fmt.Errorf("gitops blocked: TERMINATE_FORCE is forbidden in prod")
	}

	if node.Type == "DATABASE" && (action == "TERMINATE_FORCE" || action == "TERMINATE_SAFE") {
		return fmt.Errorf("gitops blocked: database termination requires manual workflow outside MVP")
	}

	cloud := strings.ToUpper(strings.TrimSpace(node.Cloud))
	if cloud != "AWS" && cloud != "AZURE" && cloud != "GCP" {
		return fmt.Errorf("gitops blocked: unsupported cloud provider %s", node.Cloud)
	}

	return nil
}

func ValidatePRRequest(diff Diff, code InfraCode, d Decision) error {
	if err := ValidateDecisionForGitOps(d); err != nil {
		return err
	}
	if err := ValidateDiffForGitOps(diff); err != nil {
		return err
	}
	if err := ValidateNodePolicyForGitOps(diff, d); err != nil {
		return err
	}
	if err := ValidateInfraCodeForGitOps(code); err != nil {
		return err
	}
	return nil
}

func normalizeActionName(action string) string {
	action = strings.TrimSpace(action)
	action = strings.TrimPrefix(action, "SAFE_")
	return action
}
