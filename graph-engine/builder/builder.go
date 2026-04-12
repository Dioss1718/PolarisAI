package builder

import (
	"fmt"

	"github.com/diya-suryawanshi/cloud/graph-engine/graph"
	"github.com/diya-suryawanshi/cloud/graph-engine/models"
)

func BuildGraph(data map[string]interface{}) (*graph.Graph, error) {
	nodesRaw, ok := data["nodes"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid input: nodes must be an array")
	}

	edgesRaw, ok := data["edges"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid input: edges must be an array")
	}

	g := graph.NewGraph()

	for i, rawNode := range nodesRaw {
		nodeMap, ok := rawNode.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("invalid input: node %d must be an object", i)
		}

		id, err := getString(nodeMap, "id")
		if err != nil {
			return nil, fmt.Errorf("node %d: %w", i, err)
		}

		nodeType, err := getString(nodeMap, "type")
		if err != nil {
			return nil, fmt.Errorf("node %d: %w", i, err)
		}

		name, err := getString(nodeMap, "name")
		if err != nil {
			return nil, fmt.Errorf("node %d: %w", i, err)
		}

		cloud, err := getString(nodeMap, "cloud_provider")
		if err != nil {
			return nil, fmt.Errorf("node %d: %w", i, err)
		}

		region, err := getString(nodeMap, "region")
		if err != nil {
			return nil, fmt.Errorf("node %d: %w", i, err)
		}

		environment, err := getString(nodeMap, "environment")
		if err != nil {
			return nil, fmt.Errorf("node %d: %w", i, err)
		}

		exposure, err := getString(nodeMap, "exposure")
		if err != nil {
			return nil, fmt.Errorf("node %d: %w", i, err)
		}

		cost, err := getFloat(nodeMap, "cost")
		if err != nil {
			return nil, fmt.Errorf("node %d: %w", i, err)
		}

		utilization, err := getFloat(nodeMap, "utilization")
		if err != nil {
			return nil, fmt.Errorf("node %d: %w", i, err)
		}

		criticality, err := getInt(nodeMap, "criticality")
		if err != nil {
			return nil, fmt.Errorf("node %d: %w", i, err)
		}

		compliance := getStringSlice(nodeMap, "compliance_flags")

		g.AddNode(models.Node{
			ID:          id,
			Type:        nodeType,
			Name:        name,
			Cloud:       cloud,
			Region:      region,
			Environment: environment,
			Cost:        cost,
			Utilization: utilization,
			Exposure:    exposure,
			Criticality: criticality,
			Compliance:  compliance,
		})
	}

	for i, rawEdge := range edgesRaw {
		edgeMap, ok := rawEdge.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("invalid input: edge %d must be an object", i)
		}

		from, err := getString(edgeMap, "from")
		if err != nil {
			return nil, fmt.Errorf("edge %d: %w", i, err)
		}

		to, err := getString(edgeMap, "to")
		if err != nil {
			return nil, fmt.Errorf("edge %d: %w", i, err)
		}

		edgeType, err := getString(edgeMap, "type")
		if err != nil {
			return nil, fmt.Errorf("edge %d: %w", i, err)
		}

		weight, err := getInt(edgeMap, "weight")
		if err != nil {
			return nil, fmt.Errorf("edge %d: %w", i, err)
		}

		if _, ok := g.Nodes[from]; !ok {
			return nil, fmt.Errorf("edge %d: unknown source node %s", i, from)
		}
		if _, ok := g.Nodes[to]; !ok {
			return nil, fmt.Errorf("edge %d: unknown target node %s", i, to)
		}

		g.AddEdge(models.Edge{
			From:   from,
			To:     to,
			Type:   edgeType,
			Weight: weight,
		})
	}

	return g, nil
}

func getString(m map[string]interface{}, key string) (string, error) {
	val, ok := m[key]
	if !ok {
		return "", fmt.Errorf("missing field %s", key)
	}

	s, ok := val.(string)
	if !ok {
		return "", fmt.Errorf("field %s must be a string", key)
	}

	return s, nil
}

func getFloat(m map[string]interface{}, key string) (float64, error) {
	val, ok := m[key]
	if !ok {
		return 0, fmt.Errorf("missing field %s", key)
	}

	f, ok := val.(float64)
	if !ok {
		return 0, fmt.Errorf("field %s must be a number", key)
	}

	return f, nil
}

func getInt(m map[string]interface{}, key string) (int, error) {
	f, err := getFloat(m, key)
	if err != nil {
		return 0, err
	}
	return int(f), nil
}

func getStringSlice(m map[string]interface{}, key string) []string {
	raw, ok := m[key].([]interface{})
	if !ok {
		return []string{}
	}

	out := []string{}
	for _, item := range raw {
		if s, ok := item.(string); ok {
			out = append(out, s)
		}
	}
	return out
}
