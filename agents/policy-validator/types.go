package policyvalidator

type InputDecision struct {
	NodeID        string
	Action        string
	CostDelta     float64
	RiskReduction float64
}

type ValidatedDecision struct {
	NodeID      string
	Action      string
	Status      string
	FinalAction string
	Score       float64
	Reason      string
}

type Policy struct {
	MaxDowntime        float64
	NoTerminateProd    bool
	NoPublicDB         bool
	EncryptionRequired bool
}

type ValidationScores struct {
	SLA        float64
	Security   float64
	Compliance float64
	Blast      float64
}
