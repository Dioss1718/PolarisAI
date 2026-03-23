package costoptimizer

type CostSignal struct {
	NodeID      string
	Type        string
	Cost        float64
	Utilization float64

	WasteRatio  float64
	GraphImpact float64
	Env         string

	Score      float64
	Confidence float64
	Reason     string
}
