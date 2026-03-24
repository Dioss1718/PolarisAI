package actiongenerator

type Action struct {
	NodeID        string
	ActionType    string
	Variant       string
	CostDelta     float64
	RiskReduction float64
	Disruption    float64
	Score         float64
}
