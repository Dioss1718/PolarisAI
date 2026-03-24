package actiongenerator

func Score(costDelta, riskReduction, disruption float64) float64 {

	return 0.5*riskReduction +
		0.3*(-costDelta) -
		0.2*disruption
}
