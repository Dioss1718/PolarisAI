package policyvalidator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type ragRequest struct {
	Query    string `json:"query"`
	NodeType string `json:"node_type"`
	Action   string `json:"action"`
}

type ragResponse struct {
	Documents []string `json:"documents"`
	Sources   []string `json:"sources"`
}

var ragHTTPClient = &http.Client{
	Timeout: 8 * time.Second,
}

func RetrievePolicyInsight(action string, nodeType string) string {
	reqBody := ragRequest{
		Query:    action + " " + nodeType + " policy SLA compliance security",
		NodeType: nodeType,
		Action:   action,
	}

	data, err := json.Marshal(reqBody)
	if err != nil {
		return "Policy insight unavailable"
	}

	resp, err := ragHTTPClient.Post(
		"http://localhost:8000/retrieve",
		"application/json",
		bytes.NewBuffer(data),
	)
	if err != nil {
		return "Policy insight unavailable"
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "Policy insight unavailable"
	}

	if resp.StatusCode != 200 {
		return "Policy insight unavailable"
	}

	var result ragResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return "Policy insight unavailable"
	}

	if len(result.Documents) == 0 {
		return "No relevant grounded policy document found"
	}

	if len(result.Sources) > 0 {
		return fmt.Sprintf("Grounded in %s", result.Sources[0])
	}

	return "Grounded policy evidence retrieved"
}
