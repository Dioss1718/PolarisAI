package tests

import (
	"testing"

	riskengine "github.com/diya-suryawanshi/cloud/agents/security-sentinel/risk-engine"
	"github.com/diya-suryawanshi/cloud/graph-engine/builder"
)

func TestRiskComputation(t *testing.T) {

	input := map[string]interface{}{
		"nodes": []interface{}{
			map[string]interface{}{
				"id":               "n1",
				"type":             "COMPUTE",
				"name":             "Node 1",
				"cloud_provider":   "AWS",
				"region":           "us-east-1",
				"environment":      "prod",
				"cost":             15.0,
				"utilization":      0.7,
				"exposure":         "PUBLIC",
				"criticality":      float64(4),
				"compliance_flags": []interface{}{},
			},
		},
		"edges": []interface{}{},
	}

	g, err := builder.BuildGraph(input)
	if err != nil {
		t.Fatalf("unexpected graph build error: %v", err)
	}

	risks := riskengine.ComputeNodeRisk(g)

	if risks["n1"] <= 0 {
		t.Errorf("expected positive risk for public node")
	}
}
