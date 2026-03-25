package builder

import (
	"github.com/diya-suryawanshi/cloud/graph-engine/graph"
	"github.com/diya-suryawanshi/cloud/graph-engine/models"
)

func BuildGraph(data map[string]interface{}) *graph.Graph {
	g := graph.NewGraph()

	nodes := data["nodes"].([]interface{})
	edges := data["edges"].([]interface{})

	for _, n := range nodes {
		nodeMap := n.(map[string]interface{})

		var compliance []string
		if flags, ok := nodeMap["compliance_flags"].([]interface{}); ok {
			for _, f := range flags {
				compliance = append(compliance, f.(string))
			}
		}

		node := models.Node{
			ID:          nodeMap["id"].(string),
			Type:        nodeMap["type"].(string),
			Name:        nodeMap["name"].(string),
			Cloud:       nodeMap["cloud_provider"].(string),
			Region:      nodeMap["region"].(string),
			Environment: nodeMap["environment"].(string),
			Cost:        nodeMap["cost"].(float64),
			Utilization: nodeMap["utilization"].(float64),
			Exposure:    nodeMap["exposure"].(string),
			Criticality: int(nodeMap["criticality"].(float64)),
			Compliance:  compliance,
		}

		g.AddNode(node)
	}

	for _, e := range edges {
		edgeMap := e.(map[string]interface{})

		edge := models.Edge{
			From:   edgeMap["from"].(string),
			To:     edgeMap["to"].(string),
			Type:   edgeMap["type"].(string),
			Weight: int(edgeMap["weight"].(float64)),
		}

		g.AddEdge(edge)
	}

	return g
}
