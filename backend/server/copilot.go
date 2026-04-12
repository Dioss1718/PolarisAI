package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/diya-suryawanshi/cloud/backend/auth"
)

type copilotRequest struct {
	Query    string `json:"query"`
	Scenario string `json:"scenario"`
	Seed     int    `json:"seed"`
}

type aiCopilotRequest struct {
	Query interface{} `json:"query"`
	State interface{} `json:"state"`
}

type copilotResponse struct {
	Answer   string   `json:"answer"`
	Grounded bool     `json:"grounded,omitempty"`
	Sources  []string `json:"sources,omitempty"`
}

var copilotHTTPClient = &http.Client{
	Timeout: 60 * time.Second,
}

func getCopilotURL() string {
	if v := os.Getenv("AI_ENGINE_URL"); v != "" {
		return strings.TrimRight(v, "/") + "/copilot"
	}
	return "http://localhost:8000/copilot"
}

func handleCopilot(w http.ResponseWriter, r *http.Request, session auth.Session) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	var req copilotRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}

	req.Query = strings.TrimSpace(req.Query)
	if req.Query == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "query is required"})
		return
	}

	scenario := strings.TrimSpace(req.Scenario)
	if scenario == "" {
		scenario = "FULL_CHAOS"
	}

	seed := req.Seed
	if seed == 0 {
		seed = 42
	}

	result := runtimeState.latestFor(scenario, seed)
	if result == nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "no pipeline state found for the requested scenario/seed. Run governance first.",
		})
		return
	}

	filtered := filterPipelineForRole(result, session.Role)

	payload := map[string]interface{}{
		"query": req.Query,
		"state": filtered,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "failed to build copilot request"})
		return
	}

	httpReq, err := http.NewRequest(http.MethodPost, getCopilotURL(), bytes.NewBuffer(body))
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "failed to build AI request"})
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := copilotHTTPClient.Do(httpReq)
	if err != nil {
		writeJSON(w, http.StatusBadGateway, map[string]string{
			"error": fmt.Sprintf("AI copilot unavailable: %v", err),
		})
		return
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		writeJSON(w, http.StatusBadGateway, map[string]string{"error": "failed to read AI copilot response"})
		return
	}

	if resp.StatusCode != http.StatusOK {
		writeJSON(w, http.StatusBadGateway, map[string]string{
			"error": fmt.Sprintf("AI copilot error: %s", string(respBody)),
		})
		return
	}

	var aiResp copilotResponse
	if err := json.Unmarshal(respBody, &aiResp); err != nil {
		writeJSON(w, http.StatusBadGateway, map[string]string{"error": "invalid AI copilot response"})
		return
	}

	if strings.TrimSpace(aiResp.Answer) == "" {
		writeJSON(w, http.StatusBadGateway, map[string]string{"error": "AI copilot returned an empty answer"})
		return
	}

	writeJSON(w, http.StatusOK, aiResp)
}
