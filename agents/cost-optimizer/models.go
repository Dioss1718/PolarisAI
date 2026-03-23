package costoptimizer

type CostInsight struct {
	NodeID      string
	CurrentCost float64
	Waste       float64
	Reason      string
}

type CostAction struct {
	ID         string
	NodeID     string
	Type       string
	DeltaCost  float64
	Confidence float64
}
