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
	fmt.Println("🚀 Starting Unified Cloud Graph Engine")

	startTime := time.Now()

	// ==============================
	// STEP 1: Fetch Simulation Data
	// ==============================
	data, err := services.FetchSimulationData()
	if err != nil {
		log.Fatalf("❌ Failed to fetch simulation data: %v", err)
	}

	// ==============================
	// STEP 2: Convert to Map
	// ==============================
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("❌ Failed to marshal data: %v", err)
	}

	var parsed map[string]interface{}
	err = json.Unmarshal(jsonData, &parsed)
	if err != nil {
		log.Fatalf("❌ Failed to unmarshal data: %v", err)
	}

	// ==============================
	// STEP 3: Build Graph G(V, E)
	// ==============================
	graph := builder.BuildGraph(parsed)

	fmt.Println("✅ Graph successfully built")
	fmt.Printf("   Nodes: %d\n", len(graph.Nodes))
	fmt.Printf("   Edges: %d\n", len(graph.Edges))

	// ==============================
	// STEP 4: Security Sentinel
	// ==============================
	fmt.Println("\n🔐 Running Security Sentinel")

	attackPaths := securitysentinel.FindAttackPaths(graph)

	if len(attackPaths) == 0 {
		fmt.Println("⚡ No attack paths detected.")
	} else {
		fmt.Println("🔥 Attack Paths Found:")

		for i, path := range attackPaths {
			risk := riskengine.CalculatePathRisk(graph, path)

			fmt.Printf("   Path %d | Risk: %.2f\n", i+1, risk)
			fmt.Printf("   → %v\n", path)
		}
	}

	// ==============================
	// STEP 5: Risk Engine
	// ==============================
	fmt.Println("\n📊 Computing Node Risk Scores")

	nodeRisks := riskengine.ComputeNodeRisk(graph)

	for id, risk := range nodeRisks {
		fmt.Printf("   Node: %-15s | Risk Score: %.2f\n", id, risk)
	}

	fmt.Println("\n💰 Running Cost Optimizer")

	insights, actions := costoptimizer.RunCostOptimizer(graph)

	for _, i := range insights {
		fmt.Printf("   Node: %s | Waste: %.2f | Reason: %s\n", i.NodeID, i.Waste, i.Reason)
	}

	fmt.Println("\n⚙️ Cost Optimization Actions")

	for _, a := range actions {
		fmt.Printf("   Action: %-12s | Node: %-15s | Savings: %.2f\n",
			a.Type, a.NodeID, a.DeltaCost)
	}

	// ==============================
	// STEP 6: Execution Metrics
	// ==============================
	duration := time.Since(startTime)

	fmt.Println("\n⏱ Execution Summary")
	fmt.Printf("   Total Time: %v\n", duration)
	fmt.Printf("   Attack Paths: %d\n", len(attackPaths))
	fmt.Printf("   Nodes Analyzed: %d\n", len(graph.Nodes))

	fmt.Println("\n✅ Graph Engine Execution Completed Successfully")
}
