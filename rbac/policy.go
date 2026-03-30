package rbac

import "strings"

func FeatureAccess(role Role, feature Feature) AccessLevel {
	policy := map[Role]map[Feature]AccessLevel{
		Admin: {
			FeatureRunGovernance:    AccessFull,
			FeatureCloudGraph:       AccessFull,
			FeatureSimulationStudio: AccessFull,
			FeatureGovernanceAction: AccessFull,
			FeatureGitOpsView:       AccessFull,
			FeatureGitOpsMerge:      AccessFull,
			FeatureExplainability:   AccessFull,
			FeatureBillShock:        AccessFull,
			FeatureFeedbackLoop:     AccessFull,
			FeatureNotifications:    AccessFull,
		},

		DevOps: {
			FeatureRunGovernance:    AccessFull,
			FeatureCloudGraph:       AccessFull,
			FeatureSimulationStudio: AccessFull,
			FeatureGovernanceAction: AccessFull,
			FeatureGitOpsView:       AccessFull,
			FeatureGitOpsMerge:      AccessNone,
			FeatureExplainability:   AccessFull,
			FeatureBillShock:        AccessFull,
			FeatureFeedbackLoop:     AccessView,
			FeatureNotifications:    AccessFull,
		},

		Security: {
			FeatureRunGovernance:    AccessFull,
			FeatureCloudGraph:       AccessFull,
			FeatureSimulationStudio: AccessNone,
			FeatureGovernanceAction: AccessFull,
			FeatureGitOpsView:       AccessFull,
			FeatureGitOpsMerge:      AccessNone,
			FeatureExplainability:   AccessFull,
			FeatureBillShock:        AccessView,
			FeatureFeedbackLoop:     AccessNone,
			FeatureNotifications:    AccessView,
		},
	}

	rolePolicy, ok := policy[role]
	if !ok {
		return AccessNone
	}

	access, ok := rolePolicy[feature]
	if !ok {
		return AccessNone
	}

	return access
}

func Permissions() map[Role][]string {
	return map[Role][]string{
		Admin: {
			"TERMINATE",
			"DOWNSIZE",
			"SECURE",
		},
		DevOps: {
			"TERMINATE",
			"DOWNSIZE",
		},
		Security: {
			"SECURE",
		},
	}
}

func AllowedActionCategories(role Role) []string {
	switch role {
	case Admin:
		return []string{"TERMINATE", "DOWNSIZE", "SECURE"}
	case DevOps:
		return []string{"TERMINATE", "DOWNSIZE"}
	case Security:
		return []string{"SECURE"}
	default:
		return []string{}
	}
}

func IsActionCategoryAllowed(role Role, action string) bool {
	for _, c := range AllowedActionCategories(role) {
		if strings.HasPrefix(strings.ToUpper(action), c) {
			return true
		}
	}
	return false
}
