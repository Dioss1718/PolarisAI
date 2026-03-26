package feedback

type Record struct {
	NodeID        string  `json:"node_id"`
	Action        string  `json:"action"`
	CostDelta     float64 `json:"cost_delta"`
	RiskReduction float64 `json:"risk_reduction"`
	Score         float64 `json:"score"`
	Reward        float64 `json:"reward"`
	Timestamp     int64   `json:"timestamp"`
}

type Summary struct {
	AvgReward float64
	Count     int
}
