package costoptimizer

func Normalize(value, max float64) float64 {
	if max == 0 {
		return 0
	}
	return value / max
}
