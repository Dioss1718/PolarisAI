package costoptimizer

import "github.com/diya-suryawanshi/cloud/graph-engine/graph"

func Analyze(g *graph.Graph) []CostSignal {
	var signals []CostSignal

	for _, node := range g.Nodes {
		graphImpact := ComputeGraphImpact(g, node.ID)

		// Zero-cost resources must still be represented in the signal layer,
		// otherwise risky IAM roles / security groups never reach candidate generation.
		wasteRatio := 0.0
		reason := ""
		score := 0.0
		confidence := 0.0

		if node.Cost > 0 {
			wasteRatio = 1 - (node.Utilization / 100.0)
			if wasteRatio < 0 {
				wasteRatio = 0
			}

			score = ComputeScore(
				node.Cost,
				wasteRatio,
				graphImpact,
				node.Environment,
				node.Type,
				node.Exposure,
			)

			confidence = ComputeConfidence(
				graphImpact,
				node.Environment,
				node.Type,
				wasteRatio,
			)

			reason = BuildReason(
				node.Type,
				node.Environment,
				node.Exposure,
				node.Cost,
				node.Utilization,
				graphImpact,
			)
		} else {
			// Keep a lightweight signal so zero-cost security-relevant assets
			// still flow into downstream candidate generation.
			reason = BuildZeroCostReason(
				node.Type,
				node.Environment,
				node.Exposure,
				graphImpact,
			)

			confidence = ComputeZeroCostConfidence(
				graphImpact,
				node.Environment,
				node.Type,
				node.Exposure,
			)

			score = ComputeZeroCostScore(
				graphImpact,
				node.Environment,
				node.Type,
				node.Exposure,
			)
		}

		signals = append(signals, CostSignal{
			NodeID:      node.ID,
			Type:        node.Type,
			Cost:        node.Cost,
			Utilization: node.Utilization,
			WasteRatio:  wasteRatio,
			GraphImpact: graphImpact,
			Env:         node.Environment,
			Score:       score,
			Confidence:  confidence,
			Reason:      reason,
		})
	}

	return signals
}

func BuildZeroCostReason(nodeType, env, exposure string, graphImpact float64) string {
	reason := "Zero-cost asset included for security remediation evaluation"

	if exposure == "PUBLIC" {
		reason += "; public exposure detected"
	}

	if env == "PROD" {
		reason += "; production sensitivity considered"
	}

	if nodeType == "IAM_ROLE" || nodeType == "SECURITY_GROUP" || nodeType == "DATABASE" {
		reason += "; security-critical resource type"
	}

	if graphImpact > 5 {
		reason += "; dependency impact is high"
	}

	return reason
}

func ComputeZeroCostConfidence(graphImpact float64, env, nodeType, exposure string) float64 {
	conf := 0.75

	if exposure == "PUBLIC" {
		conf += 0.08
	}
	if env == "PROD" {
		conf += 0.03
	}
	if nodeType == "IAM_ROLE" || nodeType == "SECURITY_GROUP" {
		conf += 0.06
	}
	if graphImpact > 5 {
		conf -= 0.05
	}

	if conf < 0.25 {
		conf = 0.25
	}
	if conf > 0.99 {
		conf = 0.99
	}

	return conf
}

func ComputeZeroCostScore(graphImpact float64, env, nodeType, exposure string) float64 {
	score := 0.20

	if exposure == "PUBLIC" {
		score += 0.20
	}
	if env == "PROD" {
		score += 0.08
	}
	if nodeType == "IAM_ROLE" || nodeType == "SECURITY_GROUP" {
		score += 0.18
	}
	if graphImpact > 5 {
		score += 0.08
	}

	if score > 1.0 {
		score = 1.0
	}

	return score
}
