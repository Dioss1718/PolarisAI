package gitops

import "fmt"

func CompareGraphs(old *Graph, new *Graph, nodeID string) Diff {
	oldNode, ok1 := old.Nodes[nodeID]
	newNode, ok2 := new.Nodes[nodeID]

	diff := Diff{
		NodeID:    nodeID,
		OldState:  map[string]interface{}{},
		NewState:  map[string]interface{}{},
		ChangeSet: []string{},
	}

	if !ok1 || !ok2 {
		if !ok1 {
			diff.ChangeSet = append(diff.ChangeSet, "Node missing in old graph")
		}
		if !ok2 {
			diff.ChangeSet = append(diff.ChangeSet, "Node missing in new graph")
		}
		return diff
	}

	if oldNode.Exposure != newNode.Exposure {
		diff.ChangeSet = append(diff.ChangeSet,
			fmt.Sprintf("Exposure changed: %s -> %s", oldNode.Exposure, newNode.Exposure))
	}

	if oldNode.Type != newNode.Type {
		diff.ChangeSet = append(diff.ChangeSet,
			fmt.Sprintf("Type changed: %s -> %s", oldNode.Type, newNode.Type))
	}

	if oldNode.Utilization != newNode.Utilization {
		diff.ChangeSet = append(diff.ChangeSet,
			fmt.Sprintf("Utilization changed: %.2f -> %.2f", oldNode.Utilization, newNode.Utilization))
	}

	if oldNode.Cost != newNode.Cost {
		diff.ChangeSet = append(diff.ChangeSet,
			fmt.Sprintf("Cost changed: %.2f -> %.2f", oldNode.Cost, newNode.Cost))
	}

	if !equalStringSlices(oldNode.Compliance, newNode.Compliance) {
		diff.ChangeSet = append(diff.ChangeSet, "Compliance updated")
	}

	diff.OldState["node"] = oldNode
	diff.NewState["node"] = newNode

	return diff
}

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
