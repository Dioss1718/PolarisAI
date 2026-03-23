package costoptimizer

import (
	"fmt"
)

func GenerateCostActions(insights []CostInsight) []CostAction {

	var actions []CostAction

	for _, insight := range insights {

		var actionType string

		// Decide action type
		if insight.Waste > insight.CurrentCost*0.6 {
			actionType = "DELETE_UNUSED"
		} else if insight.Waste > insight.CurrentCost*0.4 {
			actionType = "STOP_IDLE"
		} else {
			actionType = "RIGHTSIZE"
		}

		savings := EstimateSavings(actionType, insight.CurrentCost)

		action := CostAction{
			ID:         fmt.Sprintf("cost-%s", insight.NodeID),
			NodeID:     insight.NodeID,
			Type:       actionType,
			DeltaCost:  savings,
			Confidence: 0.8,
		}

		actions = append(actions, action)
	}

	return actions
}
