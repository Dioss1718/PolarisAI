package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type SimulationResponse struct {
	Nodes []map[string]interface{} `json:"nodes"`
	Edges []map[string]interface{} `json:"edges"`
}

func FetchSimulationData() (*SimulationResponse, error) {
	url := "http://localhost:7000/simulate/run"

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var data SimulationResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
