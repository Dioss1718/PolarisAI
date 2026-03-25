package gitops

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"os"
)

func CreatePR(code InfraCode, d Decision) PRResponse {

	token := os.Getenv("GITHUB_TOKEN")
	repo := os.Getenv("GITHUB_REPO")

	branch := "polaris/" + d.NodeID

	content := base64.StdEncoding.EncodeToString([]byte(code.Content))

	filePayload := map[string]string{
		"message": "PolarisAI Fix: " + d.Reason,
		"content": content,
		"branch":  branch,
	}

	body, _ := json.Marshal(filePayload)

	req, _ := http.NewRequest(
		"PUT",
		"https://api.github.com/repos/"+repo+"/contents/infra/"+d.NodeID+".tf",
		bytes.NewBuffer(body),
	)

	req.Header.Set("Authorization", "token "+token)

	http.DefaultClient.Do(req)

	prPayload := map[string]string{
		"title": "PolarisAI Fix: " + d.NodeID,
		"body":  d.Reason,
		"head":  branch,
		"base":  "main",
	}

	prBody, _ := json.Marshal(prPayload)

	prReq, _ := http.NewRequest(
		"POST",
		"https://api.github.com/repos/"+repo+"/pulls",
		bytes.NewBuffer(prBody),
	)

	prReq.Header.Set("Authorization", "token "+token)

	resp, _ := http.DefaultClient.Do(prReq)

	return PRResponse{
		URL:    resp.Request.URL.String(),
		Status: "CREATED",
	}
}