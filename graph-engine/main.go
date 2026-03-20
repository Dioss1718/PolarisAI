package main

import (
	"encoding/json"
	"fmt"
	"log"

	securitysentinel "github.com/diya-suryawanshi/cloud/agents/security-sentinel"
	"github.com/diya-suryawanshi/cloud/graph-engine/builder"
	"github.com/diya-suryawanshi/cloud/graph-engine/services"
)

func main() {
	fmt.Println("Starting Unified Cloud Graph Engine")

	// STEP 1: Fetch Simulation Data
	data, err := services.FetchSimulationData()
	if err != nil {
		log.Fatalf("Failed to fetch simulation data: %v", err)
	}

	// STEP 2: Convert to Map
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Failed to marshal data: %v", err)
	}

	var parsed map[string]interface{}
	err = json.Unmarshal(jsonData, &parsed)
	if err != nil {
		log.Fatalf("Failed to unmarshal data: %v", err)
	}

	// STEP 3: Build Graph G(V, E)
	graph := builder.BuildGraph(parsed)

	fmt.Println("Graph successfully built")
	fmt.Printf("Nodes: %d\n", len(graph.Nodes))
	fmt.Printf("Edges: %d\n", len(graph.Edges))

	// STEP 4: Security Sentinel
	fmt.Println("\nRunning Security Sentinel")

	attackPaths := securitysentinel.FindAttackPaths(graph)

	if len(attackPaths) == 0 {
		fmt.Println("No attack paths detected.")
	} else {
		fmt.Println("Attack Paths Found:")
		for i, path := range attackPaths {
			risk := securitysentinel.CalculatePathRisk(graph, path)
			fmt.Printf("Path %d (Risk %.2f): %v\n", i+1, risk, path)
		}
	}

	fmt.Println("\nGraph Engine Execution Completed Successfully")
}
