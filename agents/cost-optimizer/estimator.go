package costoptimizer

func EstimateCostDelta(utilization float64, cost float64) float64 {

	factor := AdaptiveSavingsFactor(utilization)

	return cost * factor
}
