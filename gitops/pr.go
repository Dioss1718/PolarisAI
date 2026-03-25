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

// This function creates a Pull Request on GitHub using generated infra code
func CreatePR(code InfraCode, d Decision, currentGraph *Graph) PRResponse {

	// Get GitHub credentials from environment
	token := os.Getenv("GITHUB_TOKEN")
	repo := os.Getenv("GITHUB_REPO")

	// If credentials missing, return failure
	if token == "" || repo == "" {
		log.Println("Missing ENV variables")
		return PRResponse{Status: "FAILED"}
	}

	// Ensure that some code is always present
	if code.Content == "" {
		code.Content = `
resource "null_resource" "` + d.NodeID + `" {
  provisioner "local-exec" {
    command = "echo default infra applied"
  }
}
`
	}

	// Create a unique branch name using nodeID and timestamp
	branch := fmt.Sprintf("polaris/%s-%d", d.NodeID, time.Now().Unix())

	// STEP 1: Get latest SHA of main branch
	url := "https://api.github.com/repos/" + repo + "/git/ref/heads/main"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/vnd.github+json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("GitHub API error:", err)
		return PRResponse{Status: "FAILED"}
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var result map[string]interface{}
	json.Unmarshal(body, &result)

	// Extract base SHA from response
	object := result["object"].(map[string]interface{})
	baseSHA := object["sha"].(string)

	// STEP 2: Create new branch from main
	branchPayload := map[string]string{
		"ref": "refs/heads/" + branch,
		"sha": baseSHA,
	}

	bBody, _ := json.Marshal(branchPayload)

	req2, _ := http.NewRequest(
		"POST",
		"https://api.github.com/repos/"+repo+"/git/refs",
		bytes.NewBuffer(bBody),
	)

	req2.Header.Set("Authorization", "Bearer "+token)
	req2.Header.Set("Content-Type", "application/json")

	http.DefaultClient.Do(req2)

	log.Println("Branch created:", branch)

	// STEP 3: Create or update Terraform file in repo

	filePath := "infra/" + d.NodeID + ".tf"

	// Check if file already exists
	getURL := fmt.Sprintf(
		"https://api.github.com/repos/%s/contents/%s?ref=%s",
		repo, filePath, branch,
	)

	reqCheck, _ := http.NewRequest("GET", getURL, nil)
	reqCheck.Header.Set("Authorization", "Bearer "+token)

	respCheck, _ := http.DefaultClient.Do(reqCheck)

	var fileSHA string

	// If file exists, get its SHA for update
	if respCheck.StatusCode == 200 {
		var existing map[string]interface{}
		bodyCheck, _ := io.ReadAll(respCheck.Body)
		json.Unmarshal(bodyCheck, &existing)
		fileSHA = existing["sha"].(string)
	}
	respCheck.Body.Close()

	// Prepare Terraform file content
	fileContent := fmt.Sprintf(`
# Terraform Infra
# Node: %s
# Time: %d

%s
`, d.NodeID, time.Now().Unix(), code.Content)

	// Encode file content to base64
	content := base64.StdEncoding.EncodeToString([]byte(fileContent))

	// Create payload for file commit
	filePayload := map[string]interface{}{
		"message": "PolarisAI update " + d.NodeID,
		"content": content,
		"branch":  branch,
	}

	// If updating existing file, include SHA
	if fileSHA != "" {
		filePayload["sha"] = fileSHA
	}

	fBody, _ := json.Marshal(filePayload)

	req3, _ := http.NewRequest(
		"PUT",
		"https://api.github.com/repos/"+repo+"/contents/"+filePath,
		bytes.NewBuffer(fBody),
	)

	req3.Header.Set("Authorization", "Bearer "+token)
	req3.Header.Set("Content-Type", "application/json")

	resp3, err := http.DefaultClient.Do(req3)
	if err != nil {
		log.Println("File commit error:", err)
		return PRResponse{Status: "FAILED"}
	}
	defer resp3.Body.Close()

	// If commit failed, return error
	if resp3.StatusCode >= 300 {
		bodyErr, _ := io.ReadAll(resp3.Body)
		log.Println("File commit failed:", string(bodyErr))
		return PRResponse{Status: "FAILED"}
	}

	log.Println("File committed successfully")

	// STEP 4: Create Pull Request
	prPayload := map[string]string{
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

	prBody, _ := json.Marshal(prPayload)

	req4, _ := http.NewRequest(
		"POST",
		"https://api.github.com/repos/"+repo+"/pulls",
		bytes.NewBuffer(prBody),
	)

	req4.Header.Set("Authorization", "Bearer "+token)
	req4.Header.Set("Content-Type", "application/json")

	prResp, err := http.DefaultClient.Do(req4)
	if err != nil {
		log.Println("PR request failed:", err)
		return PRResponse{Status: "FAILED"}
	}
	defer prResp.Body.Close()

	body4, _ := io.ReadAll(prResp.Body)

	// If PR not created successfully
	if prResp.StatusCode != 201 {
		log.Println("PR creation failed:", string(body4))
		return PRResponse{Status: "FAILED"}
	}

	// Extract PR number
	var prResult map[string]interface{}
	json.Unmarshal(body4, &prResult)

	prNumber := int(prResult["number"].(float64))

	log.Println("PR CREATED:", prNumber)

	fmt.Printf("\nOpen PR: https://github.com/%s/pull/%d\n", repo, prNumber)

	// Return PR response
	return PRResponse{
		URL:      fmt.Sprintf("https://github.com/%s/pull/%d", repo, prNumber),
		Status:   "CREATED",
		PRNumber: prNumber,
		Branch:   branch,
	}
}
