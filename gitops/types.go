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
	URL      string `json:"url"`
	Status   string `json:"status"`
	PRNumber int    `json:"prNumber"`
	Branch   string `json:"branch"`
	NodeID   string `json:"nodeId"`
	Action   string `json:"action"`
	Message  string `json:"message"`
}

type Graph = graphpkg.Graph
