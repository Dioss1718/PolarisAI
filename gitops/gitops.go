package gitops

import (
	"log"
)

func RunGitOps(
	current *Graph,
	validated []Decision,
) []PRResponse {

	var responses []PRResponse

	for _, d := range validated {

		// Step 1: Proposed Graph
		proposed := GenerateProposedGraph(current, d)

		// Step 2: Diff
		diff := CompareGraphs(current, proposed, d.NodeID)

		// Step 3: Generate Code
		code := GenerateInfraCode(diff, d)

		// Step 4: Create PR
		pr := CreatePR(code, d)

		log.Printf("[GitOps] PR created for Node=%s", d.NodeID)

		responses = append(responses, pr)
	}

	return responses
}