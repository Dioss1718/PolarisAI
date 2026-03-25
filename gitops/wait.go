package gitops

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// This function continuously checks whether a given PR is merged or not
func WaitForPRMerge(prNumber int, branch string) bool {

	// Getting GitHub token and repo name from environment variables
	token := os.Getenv("GITHUB_TOKEN")
	repo := os.Getenv("GITHUB_REPO")

	// If token or repo is missing, we cannot proceed
	if token == "" || repo == "" {
		fmt.Println("Missing GitHub ENV (TOKEN/REPO)")
		return false
	}

	// Creating GitHub API URL for the given PR
	url := fmt.Sprintf("https://api.github.com/repos/%s/pulls/%d", repo, prNumber)

	fmt.Printf("\nWaiting for PR #%d (branch: %s)\n", prNumber, branch)

	// Infinite loop to keep checking PR status
	for {

		// Creating GET request to GitHub API
		req, _ := http.NewRequest("GET", url, nil)

		// Adding authorization header using token
		req.Header.Set("Authorization", "Bearer "+token)

		// Setting API version header
		req.Header.Set("Accept", "application/vnd.github+json")

		// Sending request
		resp, err := http.DefaultClient.Do(req)

		// If error occurs, retry after delay
		if err != nil {
			fmt.Println("API error, retrying...")
			time.Sleep(5 * time.Second)
			continue
		}

		// Reading response body
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()

		// Converting JSON response into map
		var result map[string]interface{}
		json.Unmarshal(body, &result)

		// Extracting "state" and "merged" values safely
		state, _ := result["state"].(string)
		merged, _ := result["merged"].(bool)

		// Printing current PR status
		fmt.Printf("PR #%d status -> state=%s merged=%v\n", prNumber, state, merged)

		// Case 1: PR is merged successfully
		if merged {
			fmt.Printf("PR #%d MERGED\n", prNumber)
			return true
		}

		// Case 2: PR is closed but not merged (rejected)
		if state == "closed" && !merged {
			fmt.Printf("PR #%d REJECTED (closed without merge)\n", prNumber)
			return false
		}

		// Wait before checking again
		time.Sleep(5 * time.Second)
	}
}
