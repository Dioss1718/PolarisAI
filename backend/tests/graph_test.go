package tests

import (
	"testing"

	"github.com/diya-suryawanshi/cloud/graph-engine/builder"
)

func TestGraphBuild(t *testing.T) {

	input := map[string]interface{}{
		"nodes": []interface{}{
			map[string]interface{}{
				"id":               "n1",
				"type":             "COMPUTE",
				"name":             "Node 1",
				"cloud_provider":   "AWS",
				"region":           "us-east-1",
				"environment":      "prod",
				"cost":             10.0,
				"utilization":      0.6,
				"exposure":         "PUBLIC",
				"criticality":      float64(3),
				"compliance_flags": []interface{}{},
			},
			map[string]interface{}{
				"id":               "n2",
				"type":             "DATABASE",
				"name":             "Node 2",
				"cloud_provider":   "AWS",
				"region":           "us-east-1",
				"environment":      "prod",
				"cost":             20.0,
				"utilization":      0.4,
				"exposure":         "PRIVATE",
				"criticality":      float64(2),
				"compliance_flags": []interface{}{},
			},
		},
		"edges": []interface{}{},
	}

	g := builder.BuildGraph(input)

	if len(g.Nodes) != 2 {
		t.Errorf("expected 2 nodes, got %d", len(g.Nodes))
	}
}
