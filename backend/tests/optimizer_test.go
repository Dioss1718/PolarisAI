package tests

import (
	"testing"

	negotiation "github.com/diya-suryawanshi/cloud/agents/pareto-optimizer"
	"github.com/diya-suryawanshi/cloud/graph-engine/graph"
	modelspkg "github.com/diya-suryawanshi/cloud/graph-engine/models"
)

func TestParetoOptimizer(t *testing.T) {
	g := graph.NewGraph()
	g.AddNode(modelspkg.Node{
		ID:          "n1",
		Type:        "COMPUTE",
		Name:        "Node 1",
		Cloud:       "AWS",
		Region:      "us-east-1",
		Environment: "DEV",
		Cost:        50,
		Utilization: 20,
		Exposure:    "PUBLIC",
		Criticality: 5,
	})

	actions := []negotiation.Action{
		{NodeID: "n1", ActionType: "DOWNSIZE_SMALL", CostDelta: -10, RiskReduction: 5},
		{NodeID: "n1", ActionType: "SECURE_PATCH", CostDelta: -2, RiskReduction: 10},
	}

	result := negotiation.RunParetoOptimizer(g, actions, negotiation.Weights{
		RiskWeight: 0.5,
		CostWeight: 0.5,
		Penalty:    0.1,
	})

	if len(result) == 0 {
		t.Fatalf("expected optimizer to return decisions")
	}
}
