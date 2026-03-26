package gitops

type Plugin struct{}

func (p *Plugin) Run(
	graph interface{},
	decisions []Decision,
	nodeRisks map[string]float64,
) ([]PRResponse, error) {

	current := graph.(*Graph)

	_, prs := RunGitOps(current, decisions, nodeRisks)

	return prs, nil
}
