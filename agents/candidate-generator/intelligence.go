package candidategenerator

import "math"

func ComputePriority(
	cost, risk, centrality float64,
	env, nodeType, exposure string,
	wasteRatio float64,
) float64 {
	score := 0.0

	normalizedCost := math.Min(cost/250.0, 1.0)
	normalizedRisk := math.Min(risk/10.0, 1.0)
	normalizedWaste := math.Min(math.Max(wasteRatio, 0.0), 1.0)

	// Core components
	score += 0.35 * normalizedRisk
	score += 0.30 * normalizedCost
	score += 0.25 * normalizedWaste

	// Penalize aggressive prioritization for highly central nodes
	score -= 0.20 * math.Min(centrality, 1.0)

	// Context adjustments
	if env == "PROD" {
		score -= 0.08
	}

	if exposure == "PUBLIC" {
		score += 0.08
	}

	switch nodeType {
	case "DATABASE":
		score += 0.06
	case "IAM_ROLE":
		score += 0.07
	case "LOAD_BALANCER":
		score += 0.04
	}

	if score < 0 {
		score = 0
	}

	return math.Round(score*100) / 100
}

func NewCandidate(
	nodeID string,
	action string,
	cost float64,
	risk float64,
	centrality float64,
	env string,
	nodeType string,
	exposure string,
	wasteRatio float64,
) Candidate {
	return Candidate{
		NodeID:        nodeID,
		ActionType:    action,
		BaseCost:      cost,
		BaseRisk:      risk,
		Centrality:    centrality,
		Env:           env,
		PriorityScore: ComputePriority(cost, risk, centrality, env, nodeType, exposure, wasteRatio),
	}
}
