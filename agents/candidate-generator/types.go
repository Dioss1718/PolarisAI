package candidategenerator

type Candidate struct {
	NodeID        string
	ActionType    string
	BaseCost      float64
	BaseRisk      float64
	Centrality    float64
	Env           string
	PriorityScore float64
}
