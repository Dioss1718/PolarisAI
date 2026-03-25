package Forecast

import (
	"encoding/json"
	
	"fmt"
	"net/http"
)

type Result struct {
	NodeID      string  `json:"node"`
	CurrentCost float64 `json:"current_cost"`
	Forecast30  float64 `json:"forecast_30"`
	Forecast90  float64 `json:"forecast_90"`
	BillShock   bool    `json:"bill_shock"`
	ShockReason string  `json:"shock_reason"`
}

func Get(node string) (*Result, error) {
	url := fmt.Sprintf("http://localhost:5050/forecast?node=%s", node)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var r Result
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return nil, err
	}

	return &r, nil
}