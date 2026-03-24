package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	actiongenerator "github.com/diya-suryawanshi/cloud/agents/action-generator"
	candidategenerator "github.com/diya-suryawanshi/cloud/agents/candidate-generator"
	costoptimizer "github.com/diya-suryawanshi/cloud/agents/cost-optimizer"
	negotiation "github.com/diya-suryawanshi/cloud/agents/pareto-optimizer"
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

	// ==============================
	// STEP 4 & 5: PARALLEL EXECUTION
	// ==============================

	var (
		attackPaths [][]string
		nodeRisks   map[string]float64
		signals     []costoptimizer.CostSignal
	)

	var wg sync.WaitGroup
	wg.Add(2)

	// Security Sentinel
	go func() {
		defer wg.Done()

		fmt.Println("\nRunning Security Sentinel")

		attackPaths = securitysentinel.FindAttackPaths(g)

		for i, path := range attackPaths {
			risk := riskengine.CalculatePathRisk(g, path)
			fmt.Printf("Path %d | Risk: %.2f\n", i+1, risk)
			fmt.Printf("-> %v\n", path)
		}

		nodeRisks = riskengine.ComputeNodeRisk(g)

		for id, score := range nodeRisks {
			fmt.Printf("Node: %-20s | Risk: %.2f\n", id, score)
		}
	}()

	// Cost Optimizer
	go func() {
		defer wg.Done()

		fmt.Println("\nRunning Cost Optimizer")

		signals = costoptimizer.Run(g)

		for _, s := range signals {
			fmt.Printf("Node: %-20s | WasteRatio: %.2f | Score: %.2f | Confidence: %.2f\n",
				s.NodeID, s.WasteRatio, s.Score, s.Confidence)
		}
	}()

	wg.Wait()

	// ==============================
	// STEP 6: CANDIDATE GENERATOR
	// ==============================

	fmt.Println("\nRunning Candidate Generator")

	candidates := candidategenerator.GenerateCandidates(g, signals, nodeRisks)

	for _, c := range candidates {
		fmt.Printf(
			"Node: %-20s | Candidate: %-12s | BaseCost: %.2f | BaseRisk: %.2f | Centrality: %.2f | Priority: %.2f\n",
			c.NodeID,
			c.ActionType,
			c.BaseCost,
			c.BaseRisk,
			c.Centrality,
			c.PriorityScore,
		)
	}

	// ==============================
	// STEP 7: ACTION GENERATOR
	// ==============================

	fmt.Println("\nRunning Action Generator")

	actions := actiongenerator.GenerateActions(g, candidates)

	for _, a := range actions {
		fmt.Printf(
			"Node: %-20s | Action: %-12s | Variant: %-10s | CostDelta: %.2f | RiskReduction: %.2f | Disruption: %.2f | Score: %.2f\n",
			a.NodeID,
			a.ActionType,
			a.Variant,
			a.CostDelta,
			a.RiskReduction,
			a.Disruption,
			a.Score,
		)
	}

	// ==============================
	// STEP 8: CONVERT → PARETO ACTION
	// ==============================

	var paretoActions []negotiation.Action

	for _, a := range actions {
		pa := negotiation.Action{
			NodeID:        a.NodeID,
			ActionType:    a.ActionType + "_" + a.Variant,
			CostDelta:     a.CostDelta,
			RiskReduction: a.RiskReduction,
		}
		paretoActions = append(paretoActions, pa)
	}

	// ==============================
	// STEP 9: PARETO OPTIMIZER
	// ==============================

	fmt.Println("\nRunning Pareto Optimizer")

	decisions := negotiation.RunParetoOptimizer(g, paretoActions)

	for _, d := range decisions {
		fmt.Printf(
			"Node: %-20s | Final Action: %-20s | Score: %.2f\n",
			d.NodeID,
			d.Action,
			d.Score,
		)
		fmt.Printf("Reason: %s\n", d.Reason)
	}

	// ==============================
	// FINAL SUMMARY
	// ==============================

	fmt.Println("\nExecution Summary")
	fmt.Printf("Total Time    : %v\n", time.Since(startTime))
	fmt.Printf("Attack Paths  : %d\n", len(attackPaths))
	fmt.Printf("Nodes Analyzed: %d\n", len(g.Nodes))

	fmt.Println("\nExecution Completed Successfully")
}
