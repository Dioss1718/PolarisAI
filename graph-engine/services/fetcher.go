package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type SimulationResponse struct {
	Nodes              []map[string]interface{} `json:"nodes"`
	Edges              []map[string]interface{} `json:"edges"`
	ExpectedIssues     []map[string]interface{} `json:"expected_issues,omitempty"`
	Events             []map[string]interface{} `json:"events,omitempty"`
	SimulationMetadata map[string]interface{}   `json:"simulation_metadata,omitempty"`
}

func FetchSimulationData(scenario string, seed int) (*SimulationResponse, error) {
	baseURL := "http://localhost:7000/simulate/run"

	params := url.Values{}
	if scenario != "" {
		params.Set("scenario", scenario)
	}
	if seed != 0 {
		params.Set("seed", fmt.Sprintf("%d", seed))
	}

	fullURL := baseURL
	if encoded := params.Encode(); encoded != "" {
		fullURL = baseURL + "?" + encoded
	}

	resp, err := http.Get(fullURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data SimulationResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
}
