package carbon

import (
	"strings"

	modelspkg "github.com/diya-suryawanshi/cloud/graph-engine/models"
)

func mapRegion(region string) Region {
	r := strings.ToLower(region)

	switch {
	case strings.Contains(r, "india"), strings.Contains(r, "south-1"), strings.Contains(r, "central-india"):
		return RegionIndia
	case strings.Contains(r, "us-"), strings.Contains(r, "us"):
		return RegionUS
	case strings.Contains(r, "eu-"), strings.Contains(r, "europe"):
		return RegionEU
	default:
		return RegionGlobal
	}
}

func mapType(resourceType string) ResourceType {
	switch strings.ToUpper(resourceType) {
	case "COMPUTE":
		return ResourceCompute
	case "DATABASE":
		return ResourceDatabase
	case "OBJECT_STORAGE", "STORAGE":
		return ResourceStorage
	case "LOAD_BALANCER", "SECURITY_GROUP", "NETWORK", "IAM_ROLE":
		return ResourceNetwork
	default:
		return ResourceCompute
	}
}

func defaultPowerWatts(resourceType string) float64 {
	switch strings.ToUpper(resourceType) {
	case "COMPUTE":
		return 220
	case "DATABASE":
		return 300
	case "OBJECT_STORAGE":
		return 80
	case "LOAD_BALANCER":
		return 120
	case "IAM_ROLE":
		return 20
	case "SECURITY_GROUP":
		return 20
	default:
		return 100
	}
}

func defaultHours() float64 {
	return 24 * 30
}

func FromGraphNode(n modelspkg.Node) Node {
	return Node{
		ID:          n.ID,
		Type:        mapType(n.Type),
		Region:      mapRegion(n.Region),
		Utilization: n.Utilization,
		PowerWatts:  defaultPowerWatts(n.Type),
		Hours:       defaultHours(),
	}
}

func FromGraphNodes(nodes map[string]modelspkg.Node) []Node {
	out := make([]Node, 0, len(nodes))
	for _, n := range nodes {
		out = append(out, FromGraphNode(n))
	}
	return out
}
