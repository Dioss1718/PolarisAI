package tests

import (
	"testing"

	"github.com/diya-suryawanshi/cloud/carbon"
	modelspkg "github.com/diya-suryawanshi/cloud/graph-engine/models"
)

func TestCarbonCompare(t *testing.T) {
	got := carbon.Compare(100, 60)
	if got <= 0 {
		t.Fatalf("expected positive reduction, got %f", got)
	}
}

func TestCarbonCompute(t *testing.T) {
	node := carbon.Node{
		ID:          "n1",
		Type:        carbon.ResourceCompute,
		Region:      carbon.RegionIndia,
		Utilization: 50,
		PowerWatts:  200,
		Hours:       10,
	}

	value := carbon.Compute(node)
	if value <= 0 {
		t.Fatalf("expected positive carbon value, got %f", value)
	}
}

func TestFromGraphNode(t *testing.T) {
	node := modelspkg.Node{
		ID:          "aws_vm1",
		Type:        "COMPUTE",
		Region:      "ap-south-1",
		Utilization: 25,
	}

	cn := carbon.FromGraphNode(node)

	if cn.ID != "aws_vm1" {
		t.Fatalf("expected id aws_vm1, got %s", cn.ID)
	}
	if cn.Utilization != 25 {
		t.Fatalf("expected utilization 25, got %f", cn.Utilization)
	}
}
