package gitops

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func GetPRStatus(prNumber int) (string, bool, error) {
	token := os.Getenv("GITHUB_TOKEN")
	repo := os.Getenv("GITHUB_REPO")

	if token == "" || repo == "" {
		return "", false, fmt.Errorf("missing GITHUB_TOKEN or GITHUB_REPO")
	}

	url := fmt.Sprintf("https://api.github.com/repos/%s/pulls/%d", repo, prNumber)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/vnd.github+json")

	resp, err := githubHTTPClient.Do(req)
	if err != nil {
		return "", false, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode >= 300 {
		return "", false, fmt.Errorf("github pr status check failed: %s", string(body))
	}

	var result struct {
		State  string `json:"state"`
		Merged bool   `json:"merged"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", false, err
	}

	return result.State, result.Merged, nil
}

func WaitForPRMerge(prNumber int, branch string) bool {
	maxAttempts := 24
	for attempt := 0; attempt < maxAttempts; attempt++ {
		state, merged, err := GetPRStatus(prNumber)
		if err != nil {
			time.Sleep(5 * time.Second)
			continue
		}

		if merged {
			return true
		}

		if state == "closed" && !merged {
			return false
		}

		time.Sleep(5 * time.Second)
	}

	return false
}
