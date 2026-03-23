package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	costoptimizer "github.com/diya-suryawanshi/cloud/agents/cost-optimizer"
	riskengine "github.com/diya-suryawanshi/cloud/agents/security-sentinel/risk-engine"
	securitysentinel "github.com/diya-suryawanshi/cloud/agents/security-sentinel"
	"github.com/diya-suryawanshi/cloud/graph-engine/builder"
	"github.com/diya-suryawanshi/cloud/graph-engine/services"
)

func main() {
	fmt.Println("🚀 Starting Unified Cloud Graph Engine")
	startTime := time.Now()

	// STEP 1: Fetch Simulation Data
	// services.FetchSimulationData() hits synthetic-engine port 7000
	data, err := services.FetchSimulationData()
	if err != nil {
		log.Fatalf("❌ Failed to fetch simulation data: %v", err)
	}

	// STEP 2: Convert struct → map for builder
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("❌ Failed to marshal data: %v", err)
	}
	var parsed map[string]interface{}
	if err = json.Unmarshal(jsonData, &parsed); err != nil {
		log.Fatalf("❌ Failed to unmarshal data: %v", err)
	}

	// STEP 3: Build Graph G(V,E)
	// builder.BuildGraph returns *graph.Graph
	// Variable named "g" — avoids conflict with imported "graph" package
	g := builder.BuildGraph(parsed)
	fmt.Println("✅ Graph successfully built")
	fmt.Printf("   Nodes: %d\n", len(g.Nodes))
	fmt.Printf("   Edges: %d\n", len(g.Edges))

	// STEP 4: Security Sentinel — attack_path.go
	// FindAttackPaths returns [][]string (each path is []string of node IDs)
	fmt.Println("\n🔐 Running Security Sentinel")
	attackPaths := securitysentinel.FindAttackPaths(g)

	if len(attackPaths) == 0 {
		fmt.Println("⚡ No attack paths detected.")
	} else {
		fmt.Println("🔥 Attack Paths Found:")
		for i, path := range attackPaths {
			// path = []string e.g. ["sg1","aws_vm1","aws_iam_admin","aws_db1"]
			// CalculatePathRisk is in dijkstra.go — takes graph + []string path
			risk := riskengine.CalculatePathRisk(g, path)
			fmt.Printf("   Path %d | Risk: %.2f\n", i+1, risk)
			fmt.Printf("   → %v\n", path)
		}
	}

	// STEP 5: Risk Engine — estimate.go
	// ComputeNodeRisk uses:
	//   - centrality.go (ComputeCentrality) from security-sentinel
	//   - AttackerReachability from dijkstra.go
	// Returns map[nodeID]float64
	fmt.Println("\n📊 Computing Node Risk Scores")
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
		fmt.Printf("   Node: %-20s | Risk: %.2f | Level: %s\n", id, score, level)
	}

	// STEP 6: Cost Optimizer
	// RunCostOptimizer returns insights + actions
	fmt.Println("\n💰 Running Cost Optimizer")
	insights, actions := costoptimizer.RunCostOptimizer(g)

	for _, i := range insights {
		fmt.Printf("   Node: %-20s | Waste: $%.2f | Reason: %s\n", i.NodeID, i.Waste, i.Reason)
	}

	fmt.Println("\n⚙️  Cost Optimization Actions")
	for _, a := range actions {
		fmt.Printf("   Action: %-15s | Node: %-20s | Savings: $%.2f\n",
			a.Type, a.NodeID, a.DeltaCost)
	}

	// STEP 7: Summary
	fmt.Println("\n⏱  Execution Summary")
	fmt.Printf("   Total Time    : %v\n", time.Since(startTime))
	fmt.Printf("   Attack Paths  : %d\n", len(attackPaths))
	fmt.Printf("   Nodes Analyzed: %d\n", len(g.Nodes))
	fmt.Println("\n✅ Graph Engine Execution Completed Successfully")
}
