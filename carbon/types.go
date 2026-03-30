package carbon

type ResourceType string
type Region string

const (
	ResourceCompute  ResourceType = "COMPUTE"
	ResourceDatabase ResourceType = "DATABASE"
	ResourceStorage  ResourceType = "STORAGE"
	ResourceNetwork  ResourceType = "NETWORK"
)

const (
	RegionIndia  Region = "INDIA"
	RegionUS     Region = "US"
	RegionEU     Region = "EU"
	RegionGlobal Region = "GLOBAL"
)

type Node struct {
	ID          string
	Type        ResourceType
	Region      Region
	Utilization float64 // percentage [0..100]
	PowerWatts  float64 // average watts
	Hours       float64 // hours in observed window
}

type CarbonResult struct {
	NodeID string  `json:"nodeId"`
	Value  float64 `json:"value"` // grams CO2
}

type Report struct {
	Results []CarbonResult `json:"results"`
	Total   float64        `json:"total"`
	Top     []CarbonResult `json:"top"`
}
