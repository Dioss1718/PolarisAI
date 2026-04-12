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
	Scores      ValidationScores
}

type Policy struct {
	MaxDowntime        float64 `json:"max_downtime"`
	NoTerminateProd    bool    `json:"no_terminate_prod"`
	NoPublicDB         bool    `json:"no_public_db"`
	EncryptionRequired bool    `json:"encryption_required"`
}

type ValidationScores struct {
	SLA        float64
	Security   float64
	Compliance float64
	Blast      float64
}
