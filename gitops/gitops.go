package gitops

import (
	"fmt"
	"log"
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
		if len(diff.ChangeSet) == 0 {
			fmt.Println(" - No material change detected")
		}
		for _, change := range diff.ChangeSet {
			fmt.Println(" -", change)
		}

		code := GenerateInfraCode(diff, d)
		pr := CreatePR(code, d, current)

		pr.NodeID = d.NodeID
		pr.Action = d.FinalAction

		responses = append(responses, pr)

		if pr.PRNumber != 0 {
			log.Printf("[GitOps] PR created for Node=%s | PR #%d", d.NodeID, pr.PRNumber)
		}
	}

	return finalGraph, responses
}
