package costoptimizer

type CostSignal struct {
	NodeID                string
	ResourceType          string
	CurrentCost           float64
	Utilization           float64
	WasteScore            float64
	OptimizationPotential float64
	GraphImpact           float64
	ForecastCost          float64
	Confidence            float64
	Reason                string
}

type CostCandidate struct {
	NodeID     string
	ActionType string
	DeltaCost  float64
	Score      float64
}
