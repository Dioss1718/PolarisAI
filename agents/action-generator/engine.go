package actiongenerator

import (
	candidategenerator "github.com/diya-suryawanshi/cloud/agents/candidate-generator"
	"github.com/diya-suryawanshi/cloud/graph-engine/graph"
)

func GenerateActions(
	g *graph.Graph,
	candidates []candidategenerator.Candidate,
) []Action {

	var actions []Action

	for _, c := range candidates {

		variants := ExpandVariants(c.ActionType)

		for _, v := range variants {

			if !IsValid(g, c, v) {
				continue
			}

			costDelta, riskReduction, disruption :=
				Simulate(g, c, v)

			score := Score(costDelta, riskReduction, disruption)

			actions = append(actions, Action{
				NodeID:        c.NodeID,
				ActionType:    c.ActionType,
				Variant:       v,
				CostDelta:     costDelta,
				RiskReduction: riskReduction,
				Disruption:    disruption,
				Score:         score,
			})
		}
	}

	return actions
}
