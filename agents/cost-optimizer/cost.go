package costoptimizer

import "github.com/diya-suryawanshi/cloud/graph-engine/graph"

func RunCostOptimizer(g *graph.Graph) ([]CostSignal, []CostCandidate) {

	signals := AnalyzeCostSignals(g)

	candidates := GenerateCostCandidates(signals)

	return signals, candidates
}
