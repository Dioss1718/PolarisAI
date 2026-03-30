package tests

import (
	"testing"

	"github.com/diya-suryawanshi/cloud/backend/orchestrator"
)

func TestComputeAttackMetrics(t *testing.T) {
	paths := [][]string{
		{"a", "b", "c"},
		{"a", "d"},
	}

	m := orchestrator.ComputeAttackMetrics(paths)

	if m.PathCount != 2 {
		t.Fatalf("expected 2 paths, got %d", m.PathCount)
	}

	if m.ReachableNodes != 4 {
		t.Fatalf("expected 4 reachable nodes, got %d", m.ReachableNodes)
	}
}
