package carbon

func Run(nodes []Node) Report {
	results := ComputeAll(nodes)

	return Report{
		Results: results,
		Total:   Total(results),
		Top:     TopEmitters(results, 3),
	}
}
