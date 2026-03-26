package feedback

import (
	"encoding/json"
	"os"
)

type Weights struct {
	RiskWeight float64 `json:"risk_weight"`
	CostWeight float64 `json:"cost_weight"`
	Penalty    float64 `json:"penalty"`
}

const weightsFile = "backend/feedback/weights.json"

func LoadWeights() Weights {

	data, err := os.ReadFile(weightsFile)
	if err != nil {
		return Weights{
			RiskWeight: 0.5,
			CostWeight: 0.5,
			Penalty:    0.7,
		}
	}

	var w Weights
	json.Unmarshal(data, &w)

	return w
}

func SaveWeights(w Weights) {
	data, _ := json.MarshalIndent(w, "", "  ")
	os.WriteFile(weightsFile, data, 0644)
}

func UpdateWeights(s Summary) Weights {

	w := LoadWeights()

	if s.AvgReward > 0.75 {
		w.RiskWeight += 0.05
		w.CostWeight -= 0.02
		w.Penalty -= 0.02
	} else {
		w.CostWeight += 0.05
		w.RiskWeight -= 0.02
		w.Penalty += 0.02
	}

	total := w.RiskWeight + w.CostWeight
	w.RiskWeight /= total
	w.CostWeight /= total

	SaveWeights(w)

	return w
}
