package carbon

var RegionIntensity = map[Region]float64{
	RegionIndia:  700,
	RegionUS:     400,
	RegionEU:     300,
	RegionGlobal: 500,
}

var EfficiencyFactor = map[ResourceType]float64{
	ResourceCompute:  0.90,
	ResourceDatabase: 0.80,
	ResourceStorage:  0.95,
	ResourceNetwork:  0.92,
}
