package gitops

import (
	"fmt"

	riskengine "github.com/diya-suryawanshi/cloud/agents/security-sentinel/risk-engine"
)

func RunGitOps(
	current *Graph,
	validated []Decision,
	nodeRisks map[string]float64,
) (*Graph, []PRResponse) {

	var responses []PRResponse
	finalGraph := current

	fmt.Println("\nStarting GitOps Pipeline...")

	for _, d := range validated {
		proposed := GenerateProposedGraph(current, d)
		diff := CompareGraphs(current, proposed, d.NodeID)

		fmt.Println("\nDIFF DETECTED:")
		for _, change := range diff.ChangeSet {
			fmt.Println(" -", change)
		}

		code := GenerateInfraCode(diff, d)
		pr := CreatePR(code, d, current)

		responses = append(responses, pr)

		if pr.PRNumber == 0 {
			continue
		}

		fmt.Printf("\nWaiting for PR #%d (%s)\n", pr.PRNumber, pr.Branch)

		merged := WaitForPRMerge(pr.PRNumber, pr.Branch)
		if !merged {
			fmt.Println("Merge not completed, skipping...")
			continue
		}

		fmt.Println("\nApplying merged changes...")

		newGraph := proposed

		oldMetrics := EvaluateGraph(current, nodeRisks)
		newNodeRisks := riskengine.ComputeNodeRisk(newGraph)
		newMetrics := EvaluateGraph(newGraph, newNodeRisks)

		fmt.Println("\nMetrics Comparison:")
		fmt.Printf("Old Risk: %.2f\n", oldMetrics.TotalRisk)
		fmt.Printf("New Risk: %.2f\n", newMetrics.TotalRisk)

		finalGraph = SelectBestGraph(current, newGraph, newNodeRisks)

		current = finalGraph
		nodeRisks = newNodeRisks

		fmt.Println("Graph updated after merge")
	}

	return finalGraph, responses
}
