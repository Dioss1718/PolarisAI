package orchestrator

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	actiongenerator "github.com/diya-suryawanshi/cloud/agents/action-generator"
	aiexplain "github.com/diya-suryawanshi/cloud/agents/ai-explainability"
	candidategenerator "github.com/diya-suryawanshi/cloud/agents/candidate-generator"
	costoptimizer "github.com/diya-suryawanshi/cloud/agents/cost-optimizer"
	negotiation "github.com/diya-suryawanshi/cloud/agents/pareto-optimizer"
	policyvalidator "github.com/diya-suryawanshi/cloud/agents/policy-validator"
	securitysentinel "github.com/diya-suryawanshi/cloud/agents/security-sentinel"
	riskengine "github.com/diya-suryawanshi/cloud/agents/security-sentinel/risk-engine"
	forecast "github.com/diya-suryawanshi/cloud/forecast"
	gitops "github.com/diya-suryawanshi/cloud/gitops"
	pluginpkg "github.com/diya-suryawanshi/cloud/graph-engine/plugin"

	"github.com/diya-suryawanshi/cloud/graph-engine/builder"
	"github.com/diya-suryawanshi/cloud/graph-engine/services"
)

func Run() error {
	cfg := DefaultConfig()

	fmt.Println("Starting Nexus-Ops Orchestrator")
	startTime := time.Now()

	PrintSection("Fetching Simulation Data")
	data, err := services.FetchSimulationData(cfg.Scenario, cfg.Seed)
	if err != nil {
		return fmt.Errorf("failed to fetch simulation data: %w", err)
	}

	if data.SimulationMetadata != nil {
		fmt.Printf("Scenario: %v\n", data.SimulationMetadata["scenario"])
		fmt.Printf("Seed: %v\n", data.SimulationMetadata["seed"])
	}
	if len(data.ExpectedIssues) > 0 {
		fmt.Printf("Expected Issues: %d\n", len(data.ExpectedIssues))
	}
	if len(data.Events) > 0 {
		fmt.Printf("Structured Events: %d\n", len(data.Events))
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal simulation response: %w", err)
	}

	var parsed map[string]interface{}
	if err := json.Unmarshal(jsonData, &parsed); err != nil {
		return fmt.Errorf("failed to unmarshal simulation response: %w", err)
	}

	PrintSection("Building Graph")
	g := builder.BuildGraph(parsed)
	fmt.Printf("Nodes: %d\n", len(g.Nodes))
	fmt.Printf("Edges: %d\n", len(g.Edges))

	var (
		attackPaths [][]string
		nodeRisks   map[string]float64
		signals     []costoptimizer.CostSignal
	)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		PrintSection("Running Security Sentinel")
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

	go func() {
		defer wg.Done()

		PrintSection("Running Cost Optimizer")
		signals = costoptimizer.Run(g)

		for _, s := range signals {
			fmt.Printf("Node: %-20s | WasteRatio: %.2f | Score: %.2f | Confidence: %.2f\n",
				s.NodeID, s.WasteRatio, s.Score, s.Confidence)
		}
	}()

	wg.Wait()

	PrintSection("Running Candidate Generator")
	candidates := candidategenerator.GenerateCandidates(g, signals, nodeRisks)
	for _, c := range candidates {
		fmt.Printf(
			"Node: %-20s | Candidate: %-16s | Cost: %.2f | Risk: %.2f | Centrality: %.2f | Priority: %.2f\n",
			c.NodeID, c.ActionType, c.BaseCost, c.BaseRisk, c.Centrality, c.PriorityScore,
		)
	}

	PrintSection("Running Action Generator")
	actions := actiongenerator.GenerateActions(g, candidates)
	for _, a := range actions {
		fmt.Printf(
			"Node: %-20s | Action: %-16s | Variant: %-10s | CostDelta: %.2f | RiskReduction: %.2f | Score: %.2f\n",
			a.NodeID, a.ActionType, a.Variant, a.CostDelta, a.RiskReduction, a.Score,
		)
	}

	var paretoActions []negotiation.Action
	for _, a := range actions {
		paretoActions = append(paretoActions, negotiation.Action{
			NodeID:        a.NodeID,
			ActionType:    a.ActionType + "_" + a.Variant,
			CostDelta:     a.CostDelta,
			RiskReduction: a.RiskReduction,
		})
	}

	PrintSection("Running Pareto Optimizer")
	decisions := negotiation.RunParetoOptimizer(g, paretoActions)
	for _, d := range decisions {
		fmt.Printf("Node: %-20s | Final Action: %-20s | Score: %.2f\n", d.NodeID, d.Action, d.Score)
		fmt.Printf("Reason: %s\n", d.Reason)
	}

	PrintSection("Running Policy Validator")
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

		finalScore := 0.3*scores.SLA + 0.3*scores.Security + 0.2*scores.Compliance + 0.2*scores.Blast
		status := "APPROVED"
		if finalScore < 0.5 {
			status = "REJECTED"
		}

		fmt.Printf(
			"Node: %-20s | Action: %-20s | Status: %-10s | Score: %.2f\n",
			d.NodeID, d.Action, status, finalScore,
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

	PrintSection("Running Forecast")
	for _, d := range decisionsGitops {
		res, err := forecast.Get(d.NodeID)
		if err != nil {
			fmt.Printf("Node: %-20s | Forecast Error: %v\n", d.NodeID, err)
			continue
		}

		fmt.Printf(
			"Node: %-20s | Current: %.2f | F30: %.2f | F90: %.2f | Shock: %v\n",
			res.NodeID, res.CurrentCost, res.Forecast30, res.Forecast90, res.BillShock,
		)
	}

	PrintSection("Running AI Explainability Layer")
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

	PrintSection("Running GitOps")
	var prs []gitops.PRResponse
	if pluginpkg.GitOps != nil {
		prs, _ = pluginpkg.GitOps.Run(g, decisionsGitops, nodeRisks)
	}

	merged := false
	if len(prs) > 0 {
		pr := prs[0]
		if pr.PRNumber != 0 {
			fmt.Printf("\nMerge PR #%d\n", pr.PRNumber)
			merged = gitops.WaitForPRMerge(pr.PRNumber, pr.Branch)
		}
	}

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

	paths := securitysentinel.FindAttackPaths(g)

	PrintSection("Final Graph Paths")
	for i, path := range paths {
		fmt.Printf("Path %d -> %v\n", i+1, path)
	}

	PrintSection("Execution Summary")
	fmt.Printf("Total Time    : %v\n", time.Since(startTime))
	fmt.Printf("Attack Paths  : %d\n", len(paths))
	fmt.Printf("Nodes Analyzed: %d\n", len(g.Nodes))
	fmt.Println("\nExecution Completed Successfully")

	return nil
}
