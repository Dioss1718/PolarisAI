package gitops

import "github.com/diya-suryawanshi/cloud/graph-engine/graph"

type Decision struct {
	NodeID      string
	Action      string
	FinalAction string
	Score       float64
	Reason      string
}

type Diff struct {
	NodeID    string
	OldState  map[string]interface{}
	NewState  map[string]interface{}
	ChangeSet []string
}

type InfraCode struct {
	Content string
	Format  string
}

type PRRequest struct {
	Title       string
	Description string
	Branch      string
	FilePath    string
	Content     string
}

type PRResponse struct {
	URL    string
	Status string
}

type Graph = graph.Graph
