package forecast

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Result struct {
	NodeID      string  `json:"node"`
	CurrentCost float64 `json:"current_cost"`
	Forecast30  float64 `json:"forecast_30"`
	Forecast90  float64 `json:"forecast_90"`
	BillShock   bool    `json:"bill_shock"`
	ShockReason string  `json:"shock_reason"`
}

type errorResponse struct {
	Error   string `json:"error"`
	Details string `json:"details,omitempty"`
}

var httpClient = &http.Client{
	Timeout: 8 * time.Second,
}

func Get(node string) (*Result, error) {
	escapedNode := url.QueryEscape(node)
	endpoint := fmt.Sprintf("http://localhost:5050/forecast?node=%s", escapedNode)

	resp, err := httpClient.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("forecast service unreachable: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read forecast response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		var er errorResponse
		if err := json.Unmarshal(body, &er); err == nil && er.Error != "" {
			if er.Details != "" {
				return nil, fmt.Errorf("%s: %s", er.Error, er.Details)
			}
			return nil, fmt.Errorf("%s", er.Error)
		}
		return nil, fmt.Errorf("forecast service returned status %d", resp.StatusCode)
	}

	var r Result
	if err := json.Unmarshal(body, &r); err != nil {
		return nil, fmt.Errorf("failed to decode forecast response: %w", err)
	}

	return &r, nil
}
