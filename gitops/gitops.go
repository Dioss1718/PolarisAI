package gitops

import (
	"fmt"
	"log"

	riskengine "github.com/diya-suryawanshi/cloud/agents/security-sentinel/risk-engine"
)

// This function runs the complete GitOps pipeline step by step
func RunGitOps(
	current *Graph,
	validated []Decision,
	nodeRisks map[string]float64,
) (*Graph, []PRResponse) {

	// To store PR responses
	var responses []PRResponse

	// Final graph initially same as current graph
	finalGraph := current

	fmt.Println("\nStarting GitOps Pipeline...")

	// Loop through all validated decisions
	for _, d := range validated {

		// STEP 1: Create a dummy proposed graph based on decision
		proposed := GenerateProposedGraph(current, d)

		// STEP 2: Compare current graph and proposed graph
		diff := CompareGraphs(current, proposed, d.NodeID)

		fmt.Println("\nDIFF DETECTED:")
		for _, change := range diff.ChangeSet {
			fmt.Println(" -", change)
		}

		// STEP 3: Generate infrastructure code based on diff
		code := GenerateInfraCode(diff, d)

		// STEP 4: Create Pull Request using generated code
		pr := CreatePR(code, d, current)

		// Store PR response
		responses = append(responses, pr)

		log.Printf("[GitOps] PR created for Node=%s | PR #%d", d.NodeID, pr.PRNumber)

		// STEP 5: Wait for PR to be merged
		if pr.PRNumber != 0 {

			fmt.Printf("\nWaiting for PR #%d (%s)\n", pr.PRNumber, pr.Branch)

			merged := WaitForPRMerge(pr.PRNumber, pr.Branch)

			// If PR is not merged, skip further steps
			if !merged {
				fmt.Println("Merge not completed, skipping...")
				continue
			}

			// STEP 6: Apply changes after merge
			fmt.Println("\nApplying merged changes...")

			newGraph := proposed

			// Calculate metrics for old graph
			oldMetrics := EvaluateGraph(current, nodeRisks)

			// Calculate fresh risk values for new graph
			newNodeRisks := riskengine.ComputeNodeRisk(newGraph)

			// Calculate metrics for new graph
			newMetrics := EvaluateGraph(newGraph, newNodeRisks)

			fmt.Println("\nMetrics Comparison:")
			fmt.Printf("Old Risk: %.2f\n", oldMetrics.TotalRisk)
			fmt.Printf("New Risk: %.2f\n", newMetrics.TotalRisk)

			// STEP 7: Select better graph based on risk comparison
			finalGraph = SelectBestGraph(current, newGraph, newNodeRisks)

			// STEP 8: Update current graph and node risks
			current = finalGraph
			nodeRisks = newNodeRisks

			fmt.Println("Graph updated after merge")
		}
	}

	// Return final graph and all PR responses
	return finalGraph, responses
}
