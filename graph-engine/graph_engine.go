package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	costoptimizer "github.com/diya-suryawanshi/cloud/agents/cost-optimizer"
	securitysentinel "github.com/diya-suryawanshi/cloud/agents/security-sentinel"
	riskengine "github.com/diya-suryawanshi/cloud/agents/security-sentinel/risk-engine"
	"github.com/diya-suryawanshi/cloud/graph-engine/builder"
	"github.com/diya-suryawanshi/cloud/graph-engine/services"
)

func main() {
	fmt.Println("Starting Unified Cloud Graph Engine")
	startTime := time.Now()

	// STEP 1: Fetch Simulation Data
	data, err := services.FetchSimulationData()
	if err != nil {
		log.Fatalf("Failed to fetch simulation data: %v", err)
	}

	// STEP 2: Convert struct → map
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Failed to marshal data: %v", err)
	}
	var parsed map[string]interface{}
	if err = json.Unmarshal(jsonData, &parsed); err != nil {
		log.Fatalf("Failed to unmarshal data: %v", err)
	}

	// STEP 3: Build Graph
	g := builder.BuildGraph(parsed)
	fmt.Println("Graph successfully built")
	fmt.Printf("Nodes: %d\n", len(g.Nodes))
	fmt.Printf("Edges: %d\n", len(g.Edges))

	// STEP 4: Security Sentinel
	fmt.Println("\nRunning Security Sentinel")
	attackPaths := securitysentinel.FindAttackPaths(g)

	if len(attackPaths) == 0 {
		fmt.Println("No attack paths detected.")
	} else {
		fmt.Println("Attack Paths Found:")
		for i, path := range attackPaths {
			risk := riskengine.CalculatePathRisk(g, path)
			fmt.Printf("Path %d | Risk: %.2f\n", i+1, risk)
			fmt.Printf("-> %v\n", path)
		}
	}

	// STEP 5: Risk Engine
	fmt.Println("\nComputing Node Risk Scores")
	nodeRisks := riskengine.ComputeNodeRisk(g)

	for id, score := range nodeRisks {
		level := "LOW"
		if score >= 9.0 {
			level = "CRITICAL"
		} else if score >= 7.0 {
			level = "HIGH"
		} else if score >= 4.0 {
			level = "MEDIUM"
		}
		fmt.Printf("Node: %-20s | Risk: %.2f | Level: %s\n", id, score, level)
	}

	// STEP 6: Cost Optimizer (UPDATED)
	fmt.Println("\nRunning Cost Optimizer")

	signals, candidates := costoptimizer.RunCostOptimizer(g)

	for _, s := range signals {
		fmt.Printf("Node: %-20s | WasteScore: %.2f | Forecast: %.2f | Confidence: %.2f\n",
			s.NodeID, s.WasteScore, s.ForecastCost, s.Confidence)
	}

	fmt.Println("\nCost Optimization Candidates")

	for _, c := range candidates {
		fmt.Printf("Action: %-15s | Node: %-20s | DeltaCost: %.2f | Score: %.2f\n",
			c.ActionType, c.NodeID, c.DeltaCost, c.Score)
	}

	// STEP 7: Summary
	fmt.Println("\nExecution Summary")
	fmt.Printf("Total Time    : %v\n", time.Since(startTime))
	fmt.Printf("Attack Paths  : %d\n", len(attackPaths))
	fmt.Printf("Nodes Analyzed: %d\n", len(g.Nodes))

	fmt.Println("\nGraph Engine Execution Completed Successfully")
}
