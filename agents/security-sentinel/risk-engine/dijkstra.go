package riskengine

import (
	"container/heap"
	"math"

	"github.com/diya-suryawanshi/cloud/graph-engine/graph"
)

// priority queue item
type pqItem struct {
	node     string
	priority float64
	index    int
}

type priorityQueue []*pqItem

func (pq priorityQueue) Len() int { return len(pq) }
func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}
func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *priorityQueue) Push(x interface{}) {
	item := x.(*pqItem)
	item.index = len(*pq)
	*pq = append(*pq, item)
}
func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[:n-1]
	return item
}

// Dijkstra computes attacker traversal cost from a start node.
// Lower distance means easier attacker reachability.
func Dijkstra(g *graph.Graph, start string) map[string]float64 {
	dist := make(map[string]float64)

	for id := range g.Nodes {
		dist[id] = math.Inf(1)
	}
	dist[start] = 0

	pq := &priorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &pqItem{node: start, priority: 0})

	for pq.Len() > 0 {
		item := heap.Pop(pq).(*pqItem)
		current := item.node

		if item.priority > dist[current] {
			continue
		}

		for _, edge := range g.Adjacency[current] {
			weight := traversalCost(g, edge.To, float64(edge.Weight))

			if dist[current]+weight < dist[edge.To] {
				dist[edge.To] = dist[current] + weight
				heap.Push(pq, &pqItem{
					node:     edge.To,
					priority: dist[edge.To],
				})
			}
		}
	}

	return dist
}

// traversalCost keeps cost positive, but makes attacker-favorable
// nodes cheaper to traverse.
func traversalCost(g *graph.Graph, nodeID string, base float64) float64 {
	node := g.Nodes[nodeID]
	cost := base

	switch node.Exposure {
	case "PUBLIC":
		cost += 0.5
	case "INTERNAL":
		cost += 2.0
	case "PRIVATE":
		cost += 3.0
	}

	if node.Type == "SECURITY_GROUP" {
		cost += 0.5
	}

	if node.Type == "DATABASE" {
		cost += 2.0
	}

	if cost < 0.1 {
		return 0.1
	}
	return cost
}

// AttackerReachability returns a 0–10 score.
// Higher score = easier for attacker to reach from entry points.
func AttackerReachability(g *graph.Graph, nodeID string) float64 {
	minD := math.Inf(1)

	for sid, s := range g.Nodes {
		if s.Exposure != "PUBLIC" && s.Type != "SECURITY_GROUP" {
			continue
		}

		dists := Dijkstra(g, sid)
		if d, ok := dists[nodeID]; ok && d < minD {
			minD = d
		}
	}

	if math.IsInf(minD, 1) {
		return 0.0
	}

	// normalize into 0-10 range
	score := (20.0 - minD) / 20.0 * 10.0
	if score < 0 {
		return 0
	}
	if score > 10 {
		return 10
	}
	return score
}

// CalculatePathRisk computes the cumulative risk of a discovered attack path.
func CalculatePathRisk(g *graph.Graph, path []string) float64 {
	risk := 0.0

	for i, nodeID := range path {
		node := g.Nodes[nodeID]

		risk += float64(node.Criticality)

		if node.Exposure == "PUBLIC" {
			risk += 4
		}

		if node.Type == "IAM_ROLE" {
			for _, flag := range node.Compliance {
				if flag == "ADMIN_ACCESS" {
					risk += 6
					break
				}
			}
		}

		if i > 0 {
			prev := path[i-1]
			for _, edge := range g.Adjacency[prev] {
				if edge.To == nodeID {
					risk += float64(edge.Weight) * 0.5
					break
				}
			}
		}
	}

	return math.Round(risk*100) / 100
}
