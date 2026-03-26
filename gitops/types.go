package gitops

import graphpkg "github.com/diya-suryawanshi/cloud/graph-engine/graph"

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

type PRResponse struct {
	URL      string
	Status   string
	PRNumber int
	Branch   string
}

type Graph = graphpkg.Graph
