package policyvalidator

import (
	"sync"

	"github.com/diya-suryawanshi/cloud/graph-engine/graph"
)

func RunPolicyValidator(
	g *graph.Graph,
	actions []InputDecision,
	nodeRisks map[string]float64,
) []ValidatedDecision {

	policy := LoadPolicy()

	var results []ValidatedDecision
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, a := range actions {
		wg.Add(1)

		go func(action InputDecision) {
			defer wg.Done()

			node, ok := g.Nodes[action.NodeID]
			if !ok {
				mu.Lock()
				results = append(results, ValidatedDecision{
					NodeID:      action.NodeID,
					Action:      action.Action,
					Status:      "REJECTED",
					FinalAction: action.Action,
					Score:       0,
					Reason:      "Policy validation failed: node not found in graph",
					Scores: ValidationScores{
						SLA:        0,
						Security:   0,
						Compliance: 0,
						Blast:      0,
					},
				})
				mu.Unlock()
				return
			}

			inDegree := 0
			for _, edges := range g.Adjacency {
				for _, e := range edges {
					if e.To == action.NodeID {
						inDegree++
					}
				}
			}

			outDegree := len(g.Adjacency[action.NodeID])

			centrality := float64(inDegree*2+outDegree) / 10.0
			if centrality > 1.0 {
				centrality = 1.0
			}

			risk := nodeRisks[action.NodeID]

			scores := ValidateAll(
				action,
				policy,
				node.Environment,
				node.Type,
				node.Exposure,
				centrality,
				risk,
			)

			decision := ComputeFinalDecision(action, scores, node.Environment)
			decision.Scores = scores

			insight := RetrievePolicyInsight(action.Action, node.Type)
			decision.Reason = GenerateExplanation(
				action.NodeID,
				action.Action,
				decision.Score,
				insight,
			)

			Audit(decision)

			mu.Lock()
			results = append(results, decision)
			mu.Unlock()
		}(a)
	}

	wg.Wait()
	return results
}
