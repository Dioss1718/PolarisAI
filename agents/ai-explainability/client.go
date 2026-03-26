package aiexplainability

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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
	Grounded    bool     `json:"grounded,omitempty"`
	Sources     []string `json:"sources,omitempty"`
}

var httpClient = &http.Client{
	Timeout: 45 * time.Second,
}

func getExplainURL() string {
	if v := os.Getenv("AI_ENGINE_URL"); v != "" {
		return v + "/explain"
	}
	return "http://localhost:8000/explain"
}

func GetExplanation(req AIRequest) (string, error) {
	jsonData, err := json.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("marshal error: %w", err)
	}

	httpReq, err := http.NewRequest("POST", getExplainURL(), bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("request build error: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(httpReq)
	if err != nil {
		return "", fmt.Errorf("AI service unreachable: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read error: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("AI error: %s", string(body))
	}

	var result AIResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return "", fmt.Errorf("decode error: %w | raw: %s", err, string(body))
	}

	if result.Explanation == "" {
		return "", fmt.Errorf("empty explanation from AI")
	}

	return result.Explanation, nil
}
