package models

type Node struct {
	ID          string   `json:"id"`
	Type        string   `json:"type"`
	Name        string   `json:"name"`
	Cloud       string   `json:"cloud_provider"`
	Region      string   `json:"region"`
	Environment string   `json:"environment"`
	Cost        float64  `json:"cost"`
	Utilization float64  `json:"utilization"`
	Exposure    string   `json:"exposure"`
	Criticality int      `json:"criticality"`
	Compliance  []string `json:"compliance_flags"`

	// Will be computed later by agents
	RiskScore float64 `json:"risk_score"`
}
