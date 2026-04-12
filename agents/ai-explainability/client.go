package aiexplainability

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
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
		return strings.TrimRight(v, "/") + "/explain"
	}
	return "http://localhost:8000/explain"
}

func GetExplanation(req AIRequest) (AIResponse, error) {
	jsonData, err := json.Marshal(req)
	if err != nil {
		return AIResponse{}, fmt.Errorf("marshal error: %w", err)
	}

	httpReq, err := http.NewRequest(http.MethodPost, getExplainURL(), bytes.NewBuffer(jsonData))
	if err != nil {
		return AIResponse{}, fmt.Errorf("request build error: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(httpReq)
	if err != nil {
		return AIResponse{}, fmt.Errorf("AI service unreachable: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return AIResponse{}, fmt.Errorf("read error: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return AIResponse{}, fmt.Errorf("AI error: %s", string(body))
	}

	var result AIResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return AIResponse{}, fmt.Errorf("decode error: %w | raw: %s", err, string(body))
	}

	result.Explanation = strings.TrimSpace(result.Explanation)
	if result.Explanation == "" {
		return AIResponse{}, fmt.Errorf("empty explanation from AI")
	}

	if result.Sources == nil {
		result.Sources = []string{}
	}

	return result, nil
}
