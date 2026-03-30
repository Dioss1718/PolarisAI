package rbac

import "strings"

func IsAllowed(role Role, action string) bool {
	perms := Permissions()

	allowed, ok := perms[role]
	if !ok {
		return false
	}

	for _, a := range allowed {
		if strings.HasPrefix(strings.ToUpper(action), a) {
			return true
		}
	}

	return false
}

func CanAccess(role Role, feature Feature) bool {
	return FeatureAccess(role, feature) != AccessNone
}

func CanEdit(role Role, feature Feature) bool {
	return FeatureAccess(role, feature) == AccessFull
}

func CanMergeGitOps(role Role) bool {
	return role == Admin
}

func CanResetFeedback(role Role) bool {
	return role == Admin
}

func CanRunGovernance(role Role) bool {
	return FeatureAccess(role, FeatureRunGovernance) == AccessFull
}
