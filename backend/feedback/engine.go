package feedback

import "time"

func CreateRecord(nodeID, action string,
	costDelta, riskReduction, score float64,
) Record {

	reward := Evaluate(costDelta, riskReduction)

	return Record{
		NodeID:        nodeID,
		Action:        action,
		CostDelta:     costDelta,
		RiskReduction: riskReduction,
		Score:         score,
		Reward:        reward,
		Timestamp:     time.Now().Unix(),
	}
}

func Summarize(records []Record) Summary {

	if len(records) == 0 {
		return Summary{}
	}

	var total float64

	for _, r := range records {
		total += r.Reward
	}

	return Summary{
		AvgReward: total / float64(len(records)),
		Count:     len(records),
	}
}
