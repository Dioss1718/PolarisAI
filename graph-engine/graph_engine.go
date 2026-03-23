package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	costoptimizer "github.com/diya-suryawanshi/cloud/agents/cost-optimizer"
	securitysentinel "github.com/diya-suryawanshi/cloud/agents/security-sentinel"
	riskengine "github.com/diya-suryawanshi/cloud/agents/security-sentinel/risk-engine"
	"github.com/diya-suryawanshi/cloud/forecast"
	"github.com/diya-suryawanshi/cloud/graph-engine/builder"
	"github.com/diya-suryawanshi/cloud/graph-engine/services"
)

func main() {
	fmt.Println("Starting Unified Cloud Graph Engine")
	startTime := time.Now()

	data, err := services.FetchSimulationData()
	if err != nil {
		log.Fatalf("Failed to fetch simulation data: %v", err)
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Failed to marshal data: %v", err)
	}

	var parsed map[string]interface{}
	if err = json.Unmarshal(jsonData, &parsed); err != nil {
		log.Fatalf("Failed to unmarshal data: %v", err)
	}

	g := builder.BuildGraph(parsed)

	fmt.Println("Graph successfully built")
	fmt.Printf("Nodes: %d\n", len(g.Nodes))
	fmt.Printf("Edges: %d\n", len(g.Edges))

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

	fmt.Println("\nRunning Cost Optimizer")
	signals := costoptimizer.Run(g)

	for _, s := range signals {
		fmt.Printf("Node: %-20s | WasteRatio: %.2f | Score: %.2f | Confidence: %.2f\n",
			s.NodeID, s.WasteRatio, s.Score, s.Confidence)
	}

	// 🔥 FORECAST ADD (minimal + correct)
	fmt.Println("\nRunning Cost Forecast")

	nodeCosts := map[string]float64{}
	for _, s := range signals {
		nodeCosts[s.NodeID] = s.Score
	}

	forecastNodes := []string{
		"aws_vm1", "aws_db1", "azure_vm1",
		"azure_storage1", "gcp_lb1", "gcp_storage1",
	}

	forecasts := forecast.RunAllForecasts(forecastNodes, nodeCosts)

	for _, f := range forecasts {
		fmt.Printf("Node: %-20s | Now: $%.2f | 30d: $%.2f | 90d: $%.2f\n",
			f.NodeID, f.CurrentCost, f.Forecast30, f.Forecast90)
	}

	fmt.Println("\nExecution Summary")
	fmt.Printf("Total Time    : %v\n", time.Since(startTime))
	fmt.Printf("Attack Paths  : %d\n", len(attackPaths))
	fmt.Printf("Nodes Analyzed: %d\n", len(g.Nodes))
	fmt.Printf("Forecasts     : %d\n", len(forecasts))

	fmt.Println("\nGraph Engine Execution Completed Successfully")
}
