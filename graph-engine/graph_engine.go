package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/joho/godotenv"

	actiongenerator "github.com/diya-suryawanshi/cloud/agents/action-generator"
	aiexplain "github.com/diya-suryawanshi/cloud/agents/ai-explainability"
	candidategenerator "github.com/diya-suryawanshi/cloud/agents/candidate-generator"
	costoptimizer "github.com/diya-suryawanshi/cloud/agents/cost-optimizer"
	negotiation "github.com/diya-suryawanshi/cloud/agents/pareto-optimizer"
	policyvalidator "github.com/diya-suryawanshi/cloud/agents/policy-validator"
	securitysentinel "github.com/diya-suryawanshi/cloud/agents/security-sentinel"
	riskengine "github.com/diya-suryawanshi/cloud/agents/security-sentinel/risk-engine"
	gitops "github.com/diya-suryawanshi/cloud/gitops"

	pluginpkg "github.com/diya-suryawanshi/cloud/graph-engine/plugin"

	forecast "github.com/diya-suryawanshi/cloud/forecast"

	"github.com/diya-suryawanshi/cloud/graph-engine/builder"
	"github.com/diya-suryawanshi/cloud/graph-engine/services"
)

func main() {

	fmt.Println("Starting Unified Cloud Graph Engine")
	startTime := time.Now()

	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println(".env not found, using system env")
	}
	// register gitops plugin
	pluginpkg.GitOps = &gitops.Plugin{}

	// STEP 1: Fetch data
	data, err := services.FetchSimulationData()
	if err != nil {
		log.Fatalf("Failed to fetch simulation data: %v", err)
	}

	// Convert to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Failed to marshal data: %v", err)
	}

	// Parse JSON
	var parsed map[string]interface{}
	if err = json.Unmarshal(jsonData, &parsed); err != nil {
		log.Fatalf("Failed to unmarshal data: %v", err)
	}

	// STEP 2: Build graph
	g := builder.BuildGraph(parsed)

	fmt.Println("Graph successfully built")
	fmt.Printf("Nodes: %d\n", len(g.Nodes))
	fmt.Printf("Edges: %d\n", len(g.Edges))

	// STEP 3: Run security + cost in parallel
	var (
		attackPaths [][]string
		nodeRisks   map[string]float64
		signals     []costoptimizer.CostSignal
	)

	var wg sync.WaitGroup
	wg.Add(2)

	// Security analysis
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

	// Cost analysis
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

	// STEP 4: Candidate generation
	fmt.Println("\nRunning Candidate Generator")

	candidates := candidategenerator.GenerateCandidates(g, signals, nodeRisks)

	for _, c := range candidates {
		fmt.Printf(
			"Node: %-20s | Candidate: %-12s | Cost: %.2f | Risk: %.2f | Centrality: %.2f | Priority: %.2f\n",
			c.NodeID,
			c.ActionType,
			c.BaseCost,
			c.BaseRisk,
			c.Centrality,
			c.PriorityScore,
		)
	}

	// STEP 5: Action generation
	fmt.Println("\nRunning Action Generator")

	actions := actiongenerator.GenerateActions(g, candidates)

	for _, a := range actions {
		fmt.Printf(
			"Node: %-20s | Action: %-12s | Variant: %-10s | CostDelta: %.2f | RiskReduction: %.2f | Score: %.2f\n",
			a.NodeID,
			a.ActionType,
			a.Variant,
			a.CostDelta,
			a.RiskReduction,
			a.Score,
		)
	}

	// STEP 6: Convert to Pareto input
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

	// STEP 7: Pareto optimization
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

	// STEP 8: Policy validation
	fmt.Println("\nRunning Policy Validator")

	policy := policyvalidator.Policy{
		MaxDowntime:        2.0,
		NoTerminateProd:    true,
		NoPublicDB:         true,
		EncryptionRequired: true,
	}

	decisionsGitops := make([]gitops.Decision, 0)

	for _, d := range decisions {

		node := g.Nodes[d.NodeID]

		centrality := float64(len(g.Adjacency[d.NodeID])) / 5.0
		risk := nodeRisks[d.NodeID]

		input := policyvalidator.InputDecision{
			NodeID:        d.NodeID,
			Action:        d.Action,
			CostDelta:     0,
			RiskReduction: 0,
		}

		scores := policyvalidator.ValidateAll(
			input,
			policy,
			node.Environment,
			node.Type,
			node.Exposure,
			centrality,
			risk,
		)

		finalScore :=
			0.3*scores.SLA +
				0.3*scores.Security +
				0.2*scores.Compliance +
				0.2*scores.Blast

		status := "APPROVED"
		if finalScore < 0.5 {
			status = "REJECTED"
		}

		fmt.Printf(
			"Node: %-20s | Action: %-20s | Status: %-10s | Score: %.2f\n",
			d.NodeID,
			d.Action,
			status,
			finalScore,
		)

		if status == "REJECTED" {
			continue
		}

		decisionsGitops = append(decisionsGitops, gitops.Decision{
			NodeID:      d.NodeID,
			Action:      d.Action,
			FinalAction: d.Action,
			Score:       finalScore,
			Reason:      "Policy Approved",
		})
	}

	// STEP 9: Forecast (added)
	fmt.Println("\nRunning Forecast")

	for _, d := range decisionsGitops {

		res, err := forecast.Get(d.NodeID)
		if err != nil {
			fmt.Printf("Node: %-20s | Forecast Error: %v\n", d.NodeID, err)
			continue
		}

		fmt.Printf(
			"Node: %-20s | Current: %.2f | F30: %.2f | F90: %.2f | Shock: %v\n",
			res.NodeID,
			res.CurrentCost,
			res.Forecast30,
			res.Forecast90,
			res.BillShock,
		)
	}
	// ==============================
	// STEP 9: AI EXPLAINABILITY (TOP 1%)
	// ==============================
	fmt.Println("\nRunning AI Explainability Layer")

	aiStart := time.Now()

	for _, d := range decisions {

		node := g.Nodes[d.NodeID]

		centrality := float64(len(g.Adjacency[d.NodeID])) / 5.0
		risk := nodeRisks[d.NodeID]

		input := policyvalidator.InputDecision{
			NodeID:        d.NodeID,
			Action:        d.Action,
			CostDelta:     0,
			RiskReduction: 0,
		}

		scores := policyvalidator.ValidateAll(
			input,
			policy,
			node.Environment,
			node.Type,
			node.Exposure,
			centrality,
			risk,
		)

		req := aiexplain.AIRequest{
			NodeID:        d.NodeID,
			Action:        d.Action,
			Env:           node.Environment,
			NodeType:      node.Type,
			Cost:          0,
			RiskReduction: d.Score,
			SLA:           scores.SLA,
			Security:      scores.Security,
			Compliance:    scores.Compliance,
			Blast:         scores.Blast,
		}

		start := time.Now()

		explanation, err := aiexplain.GetExplanation(req)

		duration := time.Since(start)

		fmt.Println("\n--- AI REPORT ---")
		fmt.Printf("Node: %s | Action: %s\n", d.NodeID, d.Action)
		fmt.Printf("Latency: %v\n", duration)

		if err != nil {
			fmt.Println("AI Warning:", err)
		}

		fmt.Println(explanation)
	}

	fmt.Printf("\nAI Layer Time: %v\n", time.Since(aiStart))
	// STEP 10: GitOps
	fmt.Println("running gitops")

	var prs []gitops.PRResponse

	if pluginpkg.GitOps != nil {
		prs, _ = pluginpkg.GitOps.Run(g, decisionsGitops, nodeRisks)
	}

	// STEP 11: Wait for merge
	merged := false

	if len(prs) > 0 {

		pr := prs[0]

		if pr.PRNumber != 0 {

			fmt.Printf("\nMerge PR #%d\n", pr.PRNumber)

			merged = gitops.WaitForPRMerge(pr.PRNumber, pr.Branch)
		}
	}

	// STEP 12: Apply changes
	if merged {

		newGraph := gitops.GenerateFullProposedGraph(g, decisionsGitops)

		if gitops.EvaluateGraph(newGraph, nodeRisks).TotalRisk <
			gitops.EvaluateGraph(g, nodeRisks).TotalRisk {

			g = newGraph
			fmt.Println("Graph Updated")
		}

	} else {
		fmt.Println("No PR merged")
	}

	// FINAL OUTPUT
	paths := securitysentinel.FindAttackPaths(g)

	fmt.Println("\nFinal Graph (Paths):")

	for i, path := range paths {
		fmt.Printf("Path %d -> %v\n", i+1, path)
	}

	// SUMMARY
	fmt.Println("\nExecution Summary")
	fmt.Printf("Total Time    : %v\n", time.Since(startTime))
	fmt.Printf("Attack Paths  : %d\n", len(paths))
	fmt.Printf("Nodes Analyzed: %d\n", len(g.Nodes))

	fmt.Println("\nExecution Completed Successfully")
}
