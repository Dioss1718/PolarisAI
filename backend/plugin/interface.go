package plugin

import "github.com/diya-suryawanshi/cloud/gitops"

type GitOpsPlugin interface {
	Run(
		graph interface{},
		decisions []gitops.Decision,
		nodeRisks map[string]float64,
	) ([]gitops.PRResponse, error)
}
