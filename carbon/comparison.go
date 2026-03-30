package carbon

func Compare(before, after float64) float64 {
	if before <= 0 {
		return 0
	}

	reduction := (before - after) / before * 100

	if reduction < 0 {
		return 0
	}
	if reduction > 100 {
		return 100
	}
	return reduction
}

func GreenScore(before, after float64) float64 {
	if before <= 0 {
		return 100
	}

	reduction := (before - after) / before
	score := reduction * 100 * 1.5

	if score < 0 {
		return 0
	}
	if score > 100 {
		return 100
	}
	return score
}
