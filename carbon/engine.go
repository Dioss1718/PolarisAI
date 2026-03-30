package carbon

func Compute(n Node) float64 {
	intensity := RegionIntensity[n.Region]
	if intensity == 0 {
		intensity = RegionIntensity[RegionGlobal]
	}

	eff := EfficiencyFactor[n.Type]
	if eff == 0 {
		eff = 1.0
	}

	util := n.Utilization / 100.0
	if util < 0 {
		util = 0
	}
	if util > 1 {
		util = 1
	}

	energyKWh := (n.PowerWatts * n.Hours) / 1000.0
	return energyKWh * intensity * eff * util
}

func ComputeAll(nodes []Node) []CarbonResult {
	results := make([]CarbonResult, 0, len(nodes))
	for _, n := range nodes {
		results = append(results, CarbonResult{
			NodeID: n.ID,
			Value:  Compute(n),
		})
	}
	return results
}
