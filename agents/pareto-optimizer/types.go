package paretooptimizer

type Action struct {
	NodeID        string
	ActionType    string
	CostDelta     float64
	RiskReduction float64
}

type Decision struct {
	NodeID string
	Action string
	Score  float64
	Reason string
}

type Weights struct {
	RiskWeight float64
	CostWeight float64
	Penalty    float64
}
