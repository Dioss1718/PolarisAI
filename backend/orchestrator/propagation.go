package orchestrator

import (
	"fmt"
	"math"
	"sort"

	"github.com/diya-suryawanshi/cloud/graph-engine/graph"
)

func buildBlastAffectedNodes(g *graph.Graph, root string, limit int) []string {
	seen := map[string]bool{root: true}
	queue := []string{root}
	var out []string

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for _, e := range g.Adjacency[cur] {
			if !seen[e.To] {
				seen[e.To] = true
				out = append(out, e.To)
				queue = append(queue, e.To)
				if len(out) >= limit {
					return out
				}
			}
		}
	}

	return out
}

func buildEdgeInfluence(g *graph.Graph, risks map[string]float64) []EdgeInfluenceDTO {
	var out []EdgeInfluenceDTO

	for _, e := range g.Edges {
		srcRisk := risks[e.From]
		dstRisk := risks[e.To]
		influence := (srcRisk*0.58 + dstRisk*0.42) + float64(e.Weight)*0.2
		influence = math.Round(influence*100) / 100

		reason := fmt.Sprintf("%s relationship propagates risk from %s to %s.", e.Type, e.From, e.To)

		out = append(out, EdgeInfluenceDTO{
			From:      e.From,
			To:        e.To,
			Influence: influence,
			Reason:    reason,
		})
	}

	sort.Slice(out, func(i, j int) bool {
		return out[i].Influence > out[j].Influence
	})

	return out
}
