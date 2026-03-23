package costoptimizer

func EstimateSavings(actionType string, cost float64) float64 {

	switch actionType {

	case "RIGHTSIZE":
		return cost * 0.4

	case "STOP_IDLE":
		return cost * 0.7

	case "DELETE_UNUSED":
		return cost * 0.9

	case "RESERVED_INSTANCE":
		return cost * 0.3

	default:
		return 0
	}
}
