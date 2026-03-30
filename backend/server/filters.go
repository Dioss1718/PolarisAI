package server

import (
	"strings"

	"github.com/diya-suryawanshi/cloud/backend/orchestrator"
	"github.com/diya-suryawanshi/cloud/rbac"
)

func filterPipelineForRole(result *orchestrator.PipelineResult, role rbac.Role) *orchestrator.PipelineResult {
	if result == nil {
		return nil
	}

	cloned := *result

	if !rbac.CanAccess(role, rbac.FeatureGovernanceAction) {
		cloned.Recommendations = []orchestrator.RecommendationDTO{}
	} else {
		filtered := make([]orchestrator.RecommendationDTO, 0, len(result.Recommendations))
		for _, r := range result.Recommendations {
			if rbac.IsActionCategoryAllowed(role, strings.ToUpper(r.FinalAction)) {
				filtered = append(filtered, r)
			}
		}
		cloned.Recommendations = filtered
	}

	if !rbac.CanAccess(role, rbac.FeatureExplainability) {
		cloned.Explanations = []orchestrator.ExplanationDTO{}
	}

	if !rbac.CanAccess(role, rbac.FeatureBillShock) {
		cloned.Forecasts = []orchestrator.ForecastDTO{}
	}

	if !rbac.CanAccess(role, rbac.FeatureFeedbackLoop) {
		cloned.Feedback = orchestrator.FeedbackDTO{}
	}

	if !rbac.CanAccess(role, rbac.FeatureGitOpsView) {
		cloned.GitOps = orchestrator.GitOpsDTO{}
	}

	return &cloned
}
