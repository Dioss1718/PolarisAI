package aiexplainability

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type AIRequest struct {
	NodeID        string  `json:"node_id"`
	Action        string  `json:"action"`
	Env           string  `json:"env"`
	NodeType      string  `json:"node_type"`
	Cost          float64 `json:"cost"`
	RiskReduction float64 `json:"risk_reduction"`
	SLA           float64 `json:"sla"`
	Security      float64 `json:"security"`
	Compliance    float64 `json:"compliance"`
	Blast         float64 `json:"blast"`
}

type AIResponse struct {
	Explanation string   `json:"explanation"`
	Grounded    bool     `json:"grounded"`
	Sources     []string `json:"sources"`
}

var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

var explainURL = "http://localhost:8000/explain"

func GetExplanation(req AIRequest) (string, error) {
	jsonData, err := json.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("marshal error: %v", err)
	}

	httpReq, err := http.NewRequest("POST", explainURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(httpReq)
	if err != nil {
		return "", fmt.Errorf("AI service unreachable: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("AI error: %s", string(body))
	}

	var result AIResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return "", fmt.Errorf("decode error: %v | raw: %s", err, string(body))
	}

	if result.Explanation == "" {
		return "", fmt.Errorf("empty explanation from AI")
	}

	return result.Explanation, nil
}
