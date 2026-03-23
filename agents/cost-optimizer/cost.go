package costoptimizer

import (
	"github.com/diya-suryawanshi/cloud/graph-engine/graph"
)

func RunCostOptimizer(g *graph.Graph) ([]CostInsight, []CostAction) {

	// Step 1: Analyze cost inefficiencies
	insights := AnalyzeCosts(g)

	// Step 2: Generate optimization actions
	actions := GenerateCostActions(insights)

	return insights, actions
}
