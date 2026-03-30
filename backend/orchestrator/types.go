package orchestrator

type RunRequest struct {
	Scenario   string                 `json:"scenario"`
	Seed       int                    `json:"seed"`
	ManualData map[string]interface{} `json:"manualData,omitempty"`
}

type GraphNodeDTO struct {
	ID          string   `json:"id"`
	Label       string   `json:"label"`
	Type        string   `json:"type"`
	Cloud       string   `json:"cloud"`
	Region      string   `json:"region"`
	Environment string   `json:"environment"`
	Cost        float64  `json:"cost"`
	Utilization float64  `json:"utilization"`
	Exposure    string   `json:"exposure"`
	Criticality int      `json:"criticality"`
	Compliance  []string `json:"compliance"`
	Risk        float64  `json:"risk"`
	FinalAction string   `json:"finalAction,omitempty"`
	Status      string   `json:"status,omitempty"`
}

type GraphEdgeDTO struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Type   string `json:"type"`
	Weight int    `json:"weight"`
}

type RecommendationDTO struct {
	NodeID         string  `json:"nodeId"`
	Action         string  `json:"action"`
	FinalAction    string  `json:"finalAction"`
	Status         string  `json:"status"`
	Score          float64 `json:"score"`
	Reason         string  `json:"reason"`
	Risk           float64 `json:"risk"`
	Cloud          string  `json:"cloud"`
	Type           string  `json:"type"`
	Environment    string  `json:"environment"`
	CostDelta      float64 `json:"costDelta"`
	RiskReduction  float64 `json:"riskReduction"`
	Confidence     float64 `json:"confidence"`
	SafetyLevel    string  `json:"safetyLevel"`
	BlastRadius    int     `json:"blastRadius"`
	ComplianceGain float64 `json:"complianceGain"`
	GitOpsPath     string  `json:"gitOpsPath"`
	RollbackPath   string  `json:"rollbackPath"`
}

type ExplanationDTO struct {
	NodeID      string `json:"nodeId"`
	Action      string `json:"action"`
	Explanation string `json:"explanation"`
}

type ForecastDTO struct {
	NodeID      string  `json:"nodeId"`
	CurrentCost float64 `json:"currentCost"`
	Forecast30  float64 `json:"forecast30"`
	Forecast90  float64 `json:"forecast90"`
	BillShock   bool    `json:"billShock"`
	ShockReason string  `json:"shockReason"`
}

type FeedbackDTO struct {
	AvgReward  float64 `json:"avgReward"`
	Count      int     `json:"count"`
	RiskWeight float64 `json:"riskWeight"`
	CostWeight float64 `json:"costWeight"`
	Penalty    float64 `json:"penalty"`
}

type PipelineStageDTO struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

type GitOpsPRDTO struct {
	URL      string `json:"url"`
	Status   string `json:"status"`
	PRNumber int    `json:"prNumber"`
	Branch   string `json:"branch"`
	NodeID   string `json:"nodeId"`
	Action   string `json:"action"`
	Message  string `json:"message"`
}

type GitOpsDTO struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	PRs     []GitOpsPRDTO `json:"prs"`
}

type AttackMetricsDTO struct {
	PathCount      int     `json:"pathCount"`
	AvgPathLength  float64 `json:"avgPathLength"`
	ReachableNodes int     `json:"reachableNodes"`
}

type AlertDTO struct {
	Severity     string `json:"severity"`
	Title        string `json:"title"`
	Metric       string `json:"metric"`
	Reason       string `json:"reason"`
	NodeID       string `json:"nodeId,omitempty"`
	Action       string `json:"action,omitempty"`
	WorkspaceTab string `json:"workspaceTab,omitempty"`
}

type NodeIntelDTO struct {
	NodeID           string   `json:"nodeId"`
	BlastRadius      int      `json:"blastRadius"`
	RiskInfluence    float64  `json:"riskInfluence"`
	AttackPathCount  int      `json:"attackPathCount"`
	Exposed          bool     `json:"exposed"`
	Why              string   `json:"why"`
	CostRisk         float64  `json:"costRisk"`
	ComplianceBurden float64  `json:"complianceBurden"`
	AffectedNodes    []string `json:"affectedNodes"`
}

type EdgeInfluenceDTO struct {
	From      string  `json:"from"`
	To        string  `json:"to"`
	Influence float64 `json:"influence"`
	Reason    string  `json:"reason"`
}

type NegotiationAlternativeDTO struct {
	Action        string  `json:"action"`
	Score         float64 `json:"score"`
	CostDelta     float64 `json:"costDelta"`
	RiskReduction float64 `json:"riskReduction"`
	Disruption    float64 `json:"disruption"`
	Reason        string  `json:"reason"`
}

type NegotiationTraceDTO struct {
	NodeID         string                      `json:"nodeId"`
	SelectedAction string                      `json:"selectedAction"`
	SelectedScore  float64                     `json:"selectedScore"`
	WhySelected    string                      `json:"whySelected"`
	Alternatives   []NegotiationAlternativeDTO `json:"alternatives"`
}

type SummaryDTO struct {
	TotalNodes          int     `json:"totalNodes"`
	TotalEdges          int     `json:"totalEdges"`
	AttackPathCount     int     `json:"attackPathCount"`
	AvgAttackPathLength float64 `json:"avgAttackPathLength"`
	ReachableNodes      int     `json:"reachableNodes"`
	HighRiskCount       int     `json:"highRiskCount"`
	PublicExposureCount int     `json:"publicExposureCount"`
	ApprovedCount       int     `json:"approvedCount"`
	ModifiedCount       int     `json:"modifiedCount"`
	RejectedCount       int     `json:"rejectedCount"`
	UrgentCount         int     `json:"urgentCount"`
	BillShockCount      int     `json:"billShockCount"`
	CurrentTotalCost    float64 `json:"currentTotalCost"`
	Forecast30Total     float64 `json:"forecast30Total"`
	Forecast90Total     float64 `json:"forecast90Total"`
	AverageRisk         float64 `json:"averageRisk"`
	TotalRisk           float64 `json:"totalRisk"`
	CurrentCarbonTotal  float64 `json:"currentCarbonTotal"`
	ComplianceScore     float64 `json:"complianceScore"`
	CostRiskScore       float64 `json:"costRiskScore"`
}

type ProjectedSummaryDTO struct {
	ProjectedTotalCost        float64 `json:"projectedTotalCost"`
	ProjectedAttackPathCount  int     `json:"projectedAttackPathCount"`
	ProjectedPublicExposure   int     `json:"projectedPublicExposureCount"`
	ProjectedAverageRisk      float64 `json:"projectedAverageRisk"`
	ProjectedCarbonTotal      float64 `json:"projectedCarbonTotal"`
	CarbonReductionPct        float64 `json:"carbonReductionPct"`
	GreenScore                float64 `json:"greenScore"`
	ProjectedComplianceScore  float64 `json:"projectedComplianceScore"`
	ProjectedCostRiskScore    float64 `json:"projectedCostRiskScore"`
	ProjectedRiskReductionPct float64 `json:"projectedRiskReductionPct"`
}

type PipelineResult struct {
	Scenario         string                `json:"scenario"`
	Seed             int                   `json:"seed"`
	Nodes            []GraphNodeDTO        `json:"nodes"`
	Edges            []GraphEdgeDTO        `json:"edges"`
	Recommendations  []RecommendationDTO   `json:"recommendations"`
	Explanations     []ExplanationDTO      `json:"explanations"`
	Forecasts        []ForecastDTO         `json:"forecasts"`
	Feedback         FeedbackDTO           `json:"feedback"`
	GitOps           GitOpsDTO             `json:"gitops"`
	AttackPaths      [][]string            `json:"attackPaths"`
	AttackMetrics    AttackMetricsDTO      `json:"attackMetrics"`
	Summary          SummaryDTO            `json:"summary"`
	ProjectedSummary ProjectedSummaryDTO   `json:"projectedSummary"`
	Stages           []PipelineStageDTO    `json:"stages"`
	Alerts           []AlertDTO            `json:"alerts"`
	NodeIntel        []NodeIntelDTO        `json:"nodeIntel"`
	EdgeInfluence    []EdgeInfluenceDTO    `json:"edgeInfluence"`
	Negotiations     []NegotiationTraceDTO `json:"negotiations"`
	GeneratedAt      string                `json:"generatedAt"`
}
