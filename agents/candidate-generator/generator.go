package candidategenerator

import (
	"fmt"

	costoptimizer "github.com/diya-suryawanshi/cloud/agents/cost-optimizer"
	"github.com/diya-suryawanshi/cloud/graph-engine/graph"
)

func GenerateCandidates(
	g *graph.Graph,
	signals []costoptimizer.CostSignal,
	risks map[string]float64,
) []Candidate {
	var candidates []Candidate
	seen := make(map[string]bool)
	signalByNode := make(map[string]costoptimizer.CostSignal, len(signals))

	tryAdd := func(
		nodeID string,
		action string,
		cost float64,
		risk float64,
		centrality float64,
		env string,
		nodeType string,
		exposure string,
		wasteRatio float64,
	) {
		key := fmt.Sprintf("%s|%s", nodeID, action)
		if seen[key] {
			return
		}
		seen[key] = true

		candidates = append(candidates, NewCandidate(
			nodeID,
			action,
			cost,
			risk,
			centrality,
			env,
			nodeType,
			exposure,
			wasteRatio,
		))
	}

	// Pass 1: economic + mixed security candidates from optimizer signals.
	for _, s := range signals {
		signalByNode[s.NodeID] = s

		risk := risks[s.NodeID]
		node, ok := g.Nodes[s.NodeID]
		if !ok {
			continue
		}

		centrality := ComputeNodeCentrality(g, s.NodeID)
		hasEconomicSignal := s.Cost > 0

		// 1. High waste + low structural importance -> TERMINATE
		if hasEconomicSignal &&
			s.WasteRatio > 0.75 &&
			centrality < 0.45 &&
			node.Environment != "PROD" &&
			node.Type != "DATABASE" &&
			node.Type != "IAM_ROLE" &&
			node.Type != "SECURITY_GROUP" {
			tryAdd(
				s.NodeID,
				"TERMINATE",
				s.Cost,
				risk,
				centrality,
				node.Environment,
				node.Type,
				node.Exposure,
				s.WasteRatio,
			)
		}

		// 2. Moderate waste -> DOWNSIZE
		if hasEconomicSignal && s.Utilization < 40 && s.WasteRatio > 0.40 {
			tryAdd(
				s.NodeID,
				"DOWNSIZE",
				s.Cost,
				risk,
				centrality,
				node.Environment,
				node.Type,
				node.Exposure,
				s.WasteRatio,
			)
		}

		// 3. High risk -> SECURE
		if risk > 6.0 {
			tryAdd(
				s.NodeID,
				"SECURE",
				s.Cost,
				risk,
				centrality,
				node.Environment,
				node.Type,
				node.Exposure,
				s.WasteRatio,
			)
		}

		// 4. Public exposure + risk -> SECURE
		if node.Exposure == "PUBLIC" && risk > 5.5 {
			tryAdd(
				s.NodeID,
				"SECURE",
				s.Cost,
				risk,
				centrality,
				node.Environment,
				node.Type,
				node.Exposure,
				s.WasteRatio,
			)
		}

		// 5. High centrality -> prefer secure action
		if centrality > 0.65 {
			tryAdd(
				s.NodeID,
				"SECURE",
				s.Cost,
				risk,
				centrality,
				node.Environment,
				node.Type,
				node.Exposure,
				s.WasteRatio,
			)
		}

		// 6. Low-cost but risky nodes are still security-relevant
		if s.Cost < 50 && risk > 7.0 {
			tryAdd(
				s.NodeID,
				"SECURE",
				s.Cost,
				risk,
				centrality,
				node.Environment,
				node.Type,
				node.Exposure,
				s.WasteRatio,
			)
		}

		// 7. Expensive PROD compute -> downsizing candidate
		if hasEconomicSignal &&
			node.Environment == "PROD" &&
			node.Type == "COMPUTE" &&
			s.Cost > 100 &&
			s.Utilization < 50 {
			tryAdd(
				s.NodeID,
				"DOWNSIZE",
				s.Cost,
				risk,
				centrality,
				node.Environment,
				node.Type,
				node.Exposure,
				s.WasteRatio,
			)
		}
	}

	// Pass 2: security-only fallback pass across the full graph.
	// This prevents zero-cost risky assets like IAM roles and security groups
	// from being invisible just because they have little or no spend.
	for nodeID, node := range g.Nodes {
		risk := risks[nodeID]
		centrality := ComputeNodeCentrality(g, nodeID)

		s, hasSignal := signalByNode[nodeID]
		cost := node.Cost
		wasteRatio := 0.0
		if hasSignal {
			cost = s.Cost
			wasteRatio = s.WasteRatio
		}

		securityRelevantType :=
			node.Type == "IAM_ROLE" ||
				node.Type == "SECURITY_GROUP" ||
				node.Type == "DATABASE"

		needsSecurityAction :=
			risk > 6.0 ||
				(node.Exposure == "PUBLIC" && risk > 5.0) ||
				(securityRelevantType && risk > 4.5) ||
				(centrality > 0.70 && risk > 4.5)

		if needsSecurityAction {
			tryAdd(
				nodeID,
				"SECURE",
				cost,
				risk,
				centrality,
				node.Environment,
				node.Type,
				node.Exposure,
				wasteRatio,
			)
		}
	}

	return candidates
}
