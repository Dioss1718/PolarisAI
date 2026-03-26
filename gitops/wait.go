package gitops

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func WaitForPRMerge(prNumber int, branch string) bool {
	token := os.Getenv("GITHUB_TOKEN")
	repo := os.Getenv("GITHUB_REPO")

	if token == "" || repo == "" {
		fmt.Println("Missing GitHub ENV (TOKEN/REPO)")
		return false
	}

	url := fmt.Sprintf("https://api.github.com/repos/%s/pulls/%d", repo, prNumber)

	fmt.Printf("\nWaiting for PR #%d (branch: %s)\n", prNumber, branch)

	maxAttempts := 24 // ~2 minutes at 5s interval
	for attempt := 0; attempt < maxAttempts; attempt++ {
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Accept", "application/vnd.github+json")

		resp, err := githubHTTPClient.Do(req)
		if err != nil {
			fmt.Println("API error, retrying...")
			time.Sleep(5 * time.Second)
			continue
		}

		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()

		if resp.StatusCode >= 300 {
			fmt.Println("GitHub status check failed, retrying...")
			time.Sleep(5 * time.Second)
			continue
		}

		var result struct {
			State  string `json:"state"`
			Merged bool   `json:"merged"`
		}
		json.Unmarshal(body, &result)

		fmt.Printf("PR #%d status -> state=%s merged=%v\n", prNumber, result.State, result.Merged)

		if result.Merged {
			fmt.Printf("PR #%d MERGED\n", prNumber)
			return true
		}

		if result.State == "closed" && !result.Merged {
			fmt.Printf("PR #%d REJECTED (closed without merge)\n", prNumber)
			return false
		}

		time.Sleep(5 * time.Second)
	}

	fmt.Printf("PR #%d merge wait timed out\n", prNumber)
	return false
}
