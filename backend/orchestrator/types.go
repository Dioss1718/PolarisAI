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
	NodeID      string  `json:"nodeId"`
	Action      string  `json:"action"`
	FinalAction string  `json:"finalAction"`
	Status      string  `json:"status"`
	Score       float64 `json:"score"`
	Reason      string  `json:"reason"`
	Risk        float64 `json:"risk"`
	Cloud       string  `json:"cloud"`
	Type        string  `json:"type"`
	Environment string  `json:"environment"`
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

type PipelineResult struct {
	Scenario        string              `json:"scenario"`
	Seed            int                 `json:"seed"`
	Nodes           []GraphNodeDTO      `json:"nodes"`
	Edges           []GraphEdgeDTO      `json:"edges"`
	Recommendations []RecommendationDTO `json:"recommendations"`
	Explanations    []ExplanationDTO    `json:"explanations"`
	Forecasts       []ForecastDTO       `json:"forecasts"`
	Feedback        FeedbackDTO         `json:"feedback"`
	GitOps          GitOpsDTO           `json:"gitops"`
	AttackPaths     [][]string          `json:"attackPaths"`
	Stages          []PipelineStageDTO  `json:"stages"`
	GeneratedAt     string              `json:"generatedAt"`
}
