package gitops

func CompareGraphs(old *Graph, new *Graph, nodeID string) Diff {

	oldNode, ok1 := old.Nodes[nodeID]
	newNode, ok2 := new.Nodes[nodeID]

	diff := Diff{
		NodeID:    nodeID,
		OldState:  map[string]interface{}{},
		NewState:  map[string]interface{}{},
		ChangeSet: []string{},
	}

	// 🔹 Safety check
	if !ok1 || !ok2 {
		return diff
	}

	// 🔹 Exposure change
	if oldNode.Exposure != newNode.Exposure {
		diff.ChangeSet = append(diff.ChangeSet, "Exposure changed")
	}

	// 🔹 Type change
	if oldNode.Type != newNode.Type {
		diff.ChangeSet = append(diff.ChangeSet, "Type changed")
	}

	// 🔹 Utilization change
	if oldNode.Utilization != newNode.Utilization {
		diff.ChangeSet = append(diff.ChangeSet, "Utilization changed")
	}

	// 🔹 Compliance change (deep check better than length only)
	if !equalStringSlices(oldNode.Compliance, newNode.Compliance) {
		diff.ChangeSet = append(diff.ChangeSet, "Compliance updated")
	}

	// 🔹 Store states
	diff.OldState["node"] = oldNode
	diff.NewState["node"] = newNode

	return diff
}

// 🔥 Helper (important improvement)
func equalStringSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	m := make(map[string]int)

	for _, v := range a {
		m[v]++
	}

	for _, v := range b {
		if m[v] == 0 {
			return false
		}
		m[v]--
	}

	return true
}