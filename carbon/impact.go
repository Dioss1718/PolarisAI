package carbon

import "strings"

type Action string

const (
	ActionTerminate Action = "TERMINATE"
	ActionDownsize  Action = "DOWNSIZE"
	ActionSecure    Action = "SECURE"
)

func normalizeAction(action string) Action {
	action = strings.ToUpper(action)

	switch {
	case strings.Contains(action, "TERMINATE"):
		return ActionTerminate
	case strings.Contains(action, "DOWNSIZE"):
		return ActionDownsize
	case strings.Contains(action, "SECURE"):
		return ActionSecure
	default:
		return ""
	}
}

func CarbonDelta(n Node, actionStr string) float64 {
	base := Compute(n)
	utilFactor := n.Utilization / 100.0
	if utilFactor < 0 {
		utilFactor = 0
	}
	if utilFactor > 1 {
		utilFactor = 1
	}

	switch normalizeAction(actionStr) {
	case ActionTerminate:
		return -base

	case ActionDownsize:
		reductionFactor := (1 - utilFactor) * 0.8
		return -base * reductionFactor

	case ActionSecure:
		efficiencyGain := 0.05 + (0.1 * (1 - utilFactor))
		return -base * efficiencyGain

	default:
		return 0
	}
}
