package carbon

import "sort"

func TopEmitters(data []CarbonResult, n int) []CarbonResult {
	cloned := make([]CarbonResult, len(data))
	copy(cloned, data)

	sort.Slice(cloned, func(i, j int) bool {
		return cloned[i].Value > cloned[j].Value
	})

	if n > len(cloned) {
		n = len(cloned)
	}
	if n < 0 {
		n = 0
	}
	return cloned[:n]
}

func Total(data []CarbonResult) float64 {
	sum := 0.0
	for _, d := range data {
		sum += d.Value
	}
	return sum
}
