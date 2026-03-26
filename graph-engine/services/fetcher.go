package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

type SimulationResponse struct {
	Nodes              []map[string]interface{} `json:"nodes"`
	Edges              []map[string]interface{} `json:"edges"`
	SimulationMetadata map[string]interface{}   `json:"simulation_metadata"`
	ExpectedIssues     []map[string]interface{} `json:"expected_issues"`
	Events             []map[string]interface{} `json:"events"`
}

var simulationHTTPClient = &http.Client{
	Timeout: 20 * time.Second,
}

func getSimulationBaseURL() string {
	if v := os.Getenv("SIMULATION_API_URL"); v != "" {
		return v
	}

	return "http://127.0.0.1:7000"
}

func FetchSimulationData(scenario string, seed int) (*SimulationResponse, error) {
	baseURL := getSimulationBaseURL()

	u, err := url.Parse(baseURL + "/simulate/run")
	if err != nil {
		return nil, fmt.Errorf("invalid simulation base url: %w", err)
	}

	q := u.Query()
	if scenario != "" {
		q.Set("scenario", scenario)
	}
	q.Set("seed", strconv.Itoa(seed))
	u.RawQuery = q.Encode()

	resp, err := simulationHTTPClient.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading simulation response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("simulation API returned %d: %s", resp.StatusCode, string(body))
	}

	var data SimulationResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, fmt.Errorf("failed to decode simulation response: %w", err)
	}

	return &data, nil
}
