package orchestrator

// AttackMetrics holds graph-based attack analysis
type AttackMetrics struct {
	PathCount      int
	AvgPathLength  float64
	ReachableNodes int
}

// ComputeAttackMetrics calculates graph properties
func ComputeAttackMetrics(paths [][]string) AttackMetrics {

	if len(paths) == 0 {
		return AttackMetrics{}
	}

	totalLength := 0
	nodeSet := make(map[string]struct{})

	for _, p := range paths {
		totalLength += len(p)

		for _, node := range p {
			nodeSet[node] = struct{}{}
		}
	}

	avgLength := float64(totalLength) / float64(len(paths))

	return AttackMetrics{
		PathCount:      len(paths),
		AvgPathLength:  avgLength,
		ReachableNodes: len(nodeSet),
	}
}
