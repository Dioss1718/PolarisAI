package rbac

type Role string

const (
	Admin    Role = "ADMIN"
	DevOps   Role = "DEVOPS"
	Security Role = "SECURITY"
)

type Feature string

const (
	FeatureRunGovernance    Feature = "RUN_GOVERNANCE"
	FeatureCloudGraph       Feature = "CLOUD_GRAPH"
	FeatureSimulationStudio Feature = "SIMULATION_STUDIO"
	FeatureGovernanceAction Feature = "GOVERNANCE_ACTION"
	FeatureGitOpsView       Feature = "GITOPS_VIEW"
	FeatureGitOpsMerge      Feature = "GITOPS_MERGE"
	FeatureExplainability   Feature = "EXPLAINABILITY"
	FeatureBillShock        Feature = "BILL_SHOCK"
	FeatureFeedbackLoop     Feature = "FEEDBACK_LOOP"
	FeatureNotifications    Feature = "NOTIFICATIONS"
)

type AccessLevel string

const (
	AccessFull AccessLevel = "FULL"
	AccessView AccessLevel = "VIEW"
	AccessNone AccessLevel = "NONE"
)
