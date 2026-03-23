package costoptimizer

func AdaptiveSavingsFactor(utilization float64) float64 {

	if utilization < 20 {
		return 0.8
	} else if utilization < 40 {
		return 0.5
	}

	return 0.3
}
