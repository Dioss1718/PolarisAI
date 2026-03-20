package main

import (
	"encoding/json"
	"fmt"

	"github.com/diya-suryawanshi/graph-engine/builder"
	"github.com/diya-suryawanshi/graph-engine/services"
)

func main() {
	fmt.Println("Starting Unified Cloud Graph Engine...")

	data, err := services.FetchSimulationData()
	if err != nil {
		panic(err)
	}

	// Convert struct → map for builder
	jsonData, _ := json.Marshal(data)
	var parsed map[string]interface{}
	json.Unmarshal(jsonData, &parsed)

	graph := builder.BuildGraph(parsed)

	fmt.Println("Graph successfully built!")
	fmt.Printf("Nodes: %d\n", len(graph.Nodes))
	fmt.Printf("Edges: %d\n", len(graph.Edges))

	// Print sample
	for id, node := range graph.Nodes {
		fmt.Printf("Node: %s | Type: %s | Cost: %.2f\n", id, node.Type, node.Cost)
		break
	}
}
