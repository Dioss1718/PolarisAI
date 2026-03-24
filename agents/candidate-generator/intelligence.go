package candidategenerator

func ComputePriority(cost, risk, centrality float64, env string) float64 {

	score :=
		0.4*risk +
			0.3*cost -
			0.2*centrality

	if env == "PROD" {
		score -= 0.3
	}

	return score
}

func NewCandidate(nodeID, action string, cost, risk, centrality float64, env string) Candidate {
	return Candidate{
		NodeID:        nodeID,
		ActionType:    action,
		BaseCost:      cost,
		BaseRisk:      risk,
		Centrality:    centrality,
		Env:           env,
		PriorityScore: ComputePriority(cost, risk, centrality, env),
	}
}
