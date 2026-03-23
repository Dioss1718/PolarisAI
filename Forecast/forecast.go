package forecast

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const forecastServiceURL = "http://localhost:5050/forecast"

type ForecastResult struct {
	NodeID      string  `json:"node"`
	CurrentCost float64 `json:"current_cost"`
	Forecast30  float64 `json:"forecast_30"`
	Forecast90  float64 `json:"forecast_90"`
	Upper90     float64 `json:"upper_90"`
	Lower90     float64 `json:"lower_90"`
	BillShock   bool    `json:"bill_shock"`
	ShockReason string  `json:"shock_reason"`
}


func GetForecast(nodeID string, days int, currentCost float64) (*ForecastResult, error) {
	url := fmt.Sprintf(
		"%s?node=%s&days=%d&current_cost=%f",
		forecastServiceURL,
		nodeID,
		days,
		currentCost,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("forecast service unreachable: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("forecast error: %s", string(body))
	}

	var result ForecastResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse forecast: %w", err)
	}
	return &result, nil
}

func RunAllForecasts(nodeIDs []string, nodeCosts map[string]float64) []ForecastResult {
	

	var results []ForecastResult

	for _, id := range nodeIDs {
		cost, ok := nodeCosts[id]
		if !ok {
			fmt.Printf("  Forecast skipped for %s: cost not available\n", id)
			continue
		}

		result, err := GetForecast(id, 90, cost)
		if err != nil {
			fmt.Printf("  Forecast skipped for %s: %v\n", id, err)
			continue
		}

		results = append(results, *result)
	}

	return results
}