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

			node := g.Nodes[action.NodeID]

			centrality := float64(len(g.Adjacency[action.NodeID])) / 5.0
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

			decision := ComputeFinalDecision(action, scores)

			insight := RetrievePolicyInsight(action.Action)

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
