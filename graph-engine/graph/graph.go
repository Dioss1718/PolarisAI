package graph

import "github.com/diya-suryawanshi/cloud/graph-engine/models"

type Graph struct {
	Nodes map[string]models.Node
	Edges map[string][]models.Edge
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[string]models.Node),
		Edges: make(map[string][]models.Edge),
	}
}

func (g *Graph) AddNode(node models.Node) {
	g.Nodes[node.ID] = node
}

func (g *Graph) AddEdge(edge models.Edge) {
	g.Edges[edge.From] = append(g.Edges[edge.From], edge)
}

func (g *Graph) GetNeighbors(nodeID string) []models.Edge {
	return g.Edges[nodeID]
}
