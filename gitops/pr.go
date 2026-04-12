package gitops

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

var githubHTTPClient = &http.Client{
	Timeout: 12 * time.Second,
}

func CreatePR(code InfraCode, d Decision, diff Diff) PRResponse {
	if err := ValidatePRRequest(diff, code, d); err != nil {
		log.Println(err.Error())
		return PRResponse{
			Status:  "BLOCKED",
			NodeID:  d.NodeID,
			Action:  d.FinalAction,
			Message: err.Error(),
		}
	}

	token := os.Getenv("GITHUB_TOKEN")
	repo := os.Getenv("GITHUB_REPO")

	if token == "" || repo == "" {
		log.Println("Missing GITHUB_TOKEN or GITHUB_REPO")
		return PRResponse{
			Status:  "FAILED",
			NodeID:  d.NodeID,
			Action:  d.FinalAction,
			Message: "Missing GITHUB_TOKEN or GITHUB_REPO",
		}
	}

	if code.Content == "" {
		code = fallbackInfraCode(d)
	}

	branch := fmt.Sprintf("polaris/%s-%d", d.NodeID, time.Now().Unix())

	baseSHA, err := getMainBranchSHA(repo, token)
	if err != nil || baseSHA == "" {
		log.Println("Failed to fetch main branch SHA:", err)
		return PRResponse{
			Status:  "FAILED",
			NodeID:  d.NodeID,
			Action:  d.FinalAction,
			Message: "Failed to fetch main branch SHA",
		}
	}

	if err := createBranch(repo, token, branch, baseSHA); err != nil {
		log.Println("Branch creation failed:", err)
		return PRResponse{
			Status:  "FAILED",
			NodeID:  d.NodeID,
			Action:  d.FinalAction,
			Message: err.Error(),
		}
	}

	filePath := "infra/" + d.NodeID + ".tf"

	fileSHA, err := getExistingFileSHA(repo, token, filePath, branch)
	if err != nil {
		log.Println("File lookup failed:", err)
		return PRResponse{
			Status:  "FAILED",
			NodeID:  d.NodeID,
			Action:  d.FinalAction,
			Message: err.Error(),
		}
	}

	fileContent := fmt.Sprintf(`# Terraform Infra
# Node: %s
# Time: %d

%s
`, d.NodeID, time.Now().Unix(), code.Content)

	content := base64.StdEncoding.EncodeToString([]byte(fileContent))

	filePayload := map[string]interface{}{
		"message": "PolarisAI update " + d.NodeID,
		"content": content,
		"branch":  branch,
	}

	if fileSHA != "" {
		filePayload["sha"] = fileSHA
	}

	if err := putFile(repo, token, filePath, filePayload); err != nil {
		log.Println("File commit failed:", err)
		return PRResponse{
			Status:  "FAILED",
			NodeID:  d.NodeID,
			Action:  d.FinalAction,
			Message: err.Error(),
		}
	}

	prNumber, prURL, err := createPullRequest(repo, token, d, branch)
	if err != nil {
		log.Println("PR creation failed:", err)
		return PRResponse{
			Status:  "FAILED",
			NodeID:  d.NodeID,
			Action:  d.FinalAction,
			Message: err.Error(),
		}
	}

	log.Printf("[GitOps] PR created for Node=%s | PR #%d", d.NodeID, prNumber)

	return PRResponse{
		URL:      prURL,
		Status:   "CREATED",
		PRNumber: prNumber,
		Branch:   branch,
		NodeID:   d.NodeID,
		Action:   d.FinalAction,
		Message:  "PR created after validation and explicit approval",
	}
}

func getMainBranchSHA(repo, token string) (string, error) {
	url := "https://api.github.com/repos/" + repo + "/git/ref/heads/main"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/vnd.github+json")

	resp, err := githubHTTPClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("github ref fetch failed: %s", string(body))
	}

	var result struct {
		Object struct {
			SHA string `json:"sha"`
		} `json:"object"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return result.Object.SHA, nil
}

func createBranch(repo, token, branch, baseSHA string) error {
	payload := map[string]string{
		"ref": "refs/heads/" + branch,
		"sha": baseSHA,
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest(
		"POST",
		"https://api.github.com/repos/"+repo+"/git/refs",
		bytes.NewBuffer(body),
	)

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := githubHTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("branch creation failed: %s", string(b))
	}

	return nil
}

func getExistingFileSHA(repo, token, filePath, branch string) (string, error) {
	url := fmt.Sprintf(
		"https://api.github.com/repos/%s/contents/%s?ref=%s",
		repo, filePath, branch,
	)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := githubHTTPClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return "", nil
	}

	if resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("file lookup failed: %s", string(body))
	}

	var existing struct {
		SHA string `json:"sha"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&existing); err != nil {
		return "", err
	}

	return existing.SHA, nil
}

func putFile(repo, token, filePath string, payload map[string]interface{}) error {
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest(
		"PUT",
		"https://api.github.com/repos/"+repo+"/contents/"+filePath,
		bytes.NewBuffer(body),
	)

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := githubHTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("file commit failed: %s", string(b))
	}

	return nil
}

func createPullRequest(repo, token string, d Decision, branch string) (int, string, error) {
	payload := map[string]string{
		"title": "PolarisAI Fix: " + d.NodeID,
		"body": fmt.Sprintf(
			"Node: %s\nAction: %s -> %s\nScore: %.2f\nReason: %s",
			d.NodeID,
			d.Action,
			d.FinalAction,
			d.Score,
			d.Reason,
		),
		"head": branch,
		"base": "main",
	}

	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest(
		"POST",
		"https://api.github.com/repos/"+repo+"/pulls",
		bytes.NewBuffer(body),
	)

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := githubHTTPClient.Do(req)
	if err != nil {
		return 0, "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		b, _ := io.ReadAll(resp.Body)
		return 0, "", fmt.Errorf("pr creation failed: %s", string(b))
	}

	var result struct {
		Number  int    `json:"number"`
		HTMLURL string `json:"html_url"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, "", err
	}

	return result.Number, result.HTMLURL, nil
}
