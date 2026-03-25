package gitops

// This function compares old graph and new graph for a specific node
// and returns the differences between them
func CompareGraphs(old *Graph, new *Graph, nodeID string) Diff {

	// Get node from old and new graph
	oldNode, ok1 := old.Nodes[nodeID]
	newNode, ok2 := new.Nodes[nodeID]

	// Initialize diff structure
	diff := Diff{
		NodeID:    nodeID,
		OldState:  map[string]interface{}{},
		NewState:  map[string]interface{}{},
		ChangeSet: []string{},
	}

	// Safety check: if node not present in either graph, return empty diff
	if !ok1 || !ok2 {
		return diff
	}

	// Check if exposure value changed
	if oldNode.Exposure != newNode.Exposure {
		diff.ChangeSet = append(diff.ChangeSet, "Exposure changed")
	}

	// Check if node type changed
	if oldNode.Type != newNode.Type {
		diff.ChangeSet = append(diff.ChangeSet, "Type changed")
	}

	// Check if utilization changed
	if oldNode.Utilization != newNode.Utilization {
		diff.ChangeSet = append(diff.ChangeSet, "Utilization changed")
	}

	// Check if compliance list changed (deep comparison)
	if !equalStringSlices(oldNode.Compliance, newNode.Compliance) {
		diff.ChangeSet = append(diff.ChangeSet, "Compliance updated")
	}

	// Store old and new node state for reference
	diff.OldState["node"] = oldNode
	diff.NewState["node"] = newNode

	return diff
}

// This helper function checks if two string slices are equal
// It compares elements instead of just length
func equalStringSlices(a, b []string) bool {

	// If length different, they are not equal
	if len(a) != len(b) {
		return false
	}

	// Create map to count occurrences
	m := make(map[string]int)

	// Count elements of first slice
	for _, v := range a {
		m[v]++
	}

	// Reduce count using second slice
	for _, v := range b {
		if m[v] == 0 {
			return false
		}
		m[v]--
	}

	// If all matched, slices are equal
	return true
}
