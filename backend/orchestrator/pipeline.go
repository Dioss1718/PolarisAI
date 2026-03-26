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
	feedback "github.com/diya-suryawanshi/cloud/backend/feedback"
	pluginpkg "github.com/diya-suryawanshi/cloud/backend/plugin"
	forecast "github.com/diya-suryawanshi/cloud/forecast"
	gitops "github.com/diya-suryawanshi/cloud/gitops"
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

	PrintSection("Building Unified Cloud Graph")

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
			fmt.Printf(
				"Node: %-20s | WasteRatio: %.2f | Score: %.2f | Confidence: %.2f\n",
				s.NodeID, s.WasteRatio, s.Score, s.Confidence,
			)
		}
	}()

	wg.Wait()

	PrintSection("Running Candidate Generator")

	candidates := candidategenerator.GenerateCandidates(g, signals, nodeRisks)

	for _, c := range candidates {
		fmt.Printf(
			"Node: %-20s | Candidate: %-16s | Cost: %.2f | Risk: %.2f | Centrality: %.2f | Priority: %.2f\n",
			c.NodeID,
			c.ActionType,
			c.BaseCost,
			c.BaseRisk,
			c.Centrality,
			c.PriorityScore,
		)
	}

	PrintSection("Running Action Generator")

	actions := actiongenerator.GenerateActions(g, candidates)

	for _, a := range actions {
		fmt.Printf(
			"Node: %-20s | Action: %-16s | Variant: %-10s | CostDelta: %.2f | RiskReduction: %.2f | Score: %.2f\n",
			a.NodeID,
			a.ActionType,
			a.Variant,
			a.CostDelta,
			a.RiskReduction,
			a.Score,
		)
	}

	actionLookup := make(map[string]actiongenerator.Action)
	var paretoActions []negotiation.Action

	for _, a := range actions {
		paretoActionName := a.ActionType + "_" + a.Variant

		actionLookup[a.NodeID+"|"+paretoActionName] = a

		paretoActions = append(paretoActions, negotiation.Action{
			NodeID:        a.NodeID,
			ActionType:    paretoActionName,
			CostDelta:     a.CostDelta,
			RiskReduction: a.RiskReduction,
		})
	}

	PrintSection("Running Pareto Optimizer")

	learnedWeights := feedback.LoadWeights()

	decisions := negotiation.RunParetoOptimizer(
		g,
		paretoActions,
		negotiation.Weights{
			RiskWeight: learnedWeights.RiskWeight,
			CostWeight: learnedWeights.CostWeight,
			Penalty:    learnedWeights.Penalty,
		},
	)

	decisionLookup := make(map[string]negotiation.Decision)

	for _, d := range decisions {
		decisionLookup[d.NodeID+"|"+d.Action] = d

		fmt.Printf(
			"Node: %-20s | Final Action: %-20s | Score: %.2f\n",
			d.NodeID,
			d.Action,
			d.Score,
		)
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
	approvedActions := make([]actiongenerator.Action, 0)

	for _, d := range decisions {
		node := g.Nodes[d.NodeID]

		inDegree := 0
		for _, edges := range g.Adjacency {
			for _, e := range edges {
				if e.To == d.NodeID {
					inDegree++
				}
			}
		}
		outDegree := len(g.Adjacency[d.NodeID])
		centrality := float64(inDegree*2+outDegree) / 10.0
		if centrality > 1.0 {
			centrality = 1.0
		}

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
		finalAction := d.Action

		if finalScore < 0.4 {
			status = "REJECTED"
		} else if finalScore < 0.7 {
			status = "MODIFIED"
			finalAction = "SAFE_" + d.Action
		}

		fmt.Printf(
			"Node: %-20s | Action: %-20s | Status: %-10s | Score: %.2f\n",
			d.NodeID,
			finalAction,
			status,
			finalScore,
		)

		if status == "REJECTED" {
			continue
		}

		decisionsGitops = append(decisionsGitops, gitops.Decision{
			NodeID:      d.NodeID,
			Action:      d.Action,
			FinalAction: finalAction,
			Score:       finalScore,
			Reason:      "Policy Approved",
		})

		if srcAction, ok := actionLookup[d.NodeID+"|"+d.Action]; ok {
			approvedActions = append(approvedActions, srcAction)
		}
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
			res.NodeID,
			res.CurrentCost,
			res.Forecast30,
			res.Forecast90,
			res.BillShock,
		)
	}

	PrintSection("Running AI Explainability Layer")
	aiStart := time.Now()

	for _, d := range decisionsGitops {
		node := g.Nodes[d.NodeID]

		inDegree := 0
		for _, edges := range g.Adjacency {
			for _, e := range edges {
				if e.To == d.NodeID {
					inDegree++
				}
			}
		}
		outDegree := len(g.Adjacency[d.NodeID])
		centrality := float64(inDegree*2+outDegree) / 10.0
		if centrality > 1.0 {
			centrality = 1.0
		}

		risk := nodeRisks[d.NodeID]

		input := policyvalidator.InputDecision{
			NodeID:        d.NodeID,
			Action:        d.FinalAction,
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

		originalDecisionScore := d.Score
		if original, ok := decisionLookup[d.NodeID+"|"+d.Action]; ok {
			originalDecisionScore = original.Score
		}

		req := aiexplain.AIRequest{
			NodeID:        d.NodeID,
			Action:        d.FinalAction,
			Env:           node.Environment,
			NodeType:      node.Type,
			Cost:          0,
			RiskReduction: originalDecisionScore,
			SLA:           scores.SLA,
			Security:      scores.Security,
			Compliance:    scores.Compliance,
			Blast:         scores.Blast,
		}

		start := time.Now()
		explanation, err := aiexplain.GetExplanation(req)
		duration := time.Since(start)

		fmt.Println("\n--- AI REPORT ---")
		fmt.Printf("Node: %s | Action: %s\n", d.NodeID, d.FinalAction)
		fmt.Printf("Latency: %v\n", duration)

		if err != nil {
			fmt.Println("AI Warning:", err)
		} else {
			fmt.Println(explanation)
		}
	}

	fmt.Printf("\nAI Layer Time: %v\n", time.Since(aiStart))

	PrintSection("Running GitOps")

	var prs []gitops.PRResponse
	if pluginpkg.GitOps != nil {
		var gitopsErr error
		prs, gitopsErr = pluginpkg.GitOps.Run(g, decisionsGitops, nodeRisks)
		if gitopsErr != nil {
			fmt.Printf("GitOps skipped: %v\n", gitopsErr)
		}
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

	PrintSection("Running Feedback Learning Loop")

	records := feedback.Load()

	for _, a := range approvedActions {
		rec := feedback.CreateRecord(
			a.NodeID,
			a.ActionType+"_"+a.Variant,
			a.CostDelta,
			a.RiskReduction,
			a.Score,
		)
		records = append(records, rec)
	}

	feedback.Save(records)

	summary := feedback.Summarize(records)
	newWeights := feedback.UpdateWeights(summary)

	fmt.Printf("Feedback Avg Reward: %.2f\n", summary.AvgReward)
	fmt.Printf("Feedback Count: %d\n", summary.Count)
	fmt.Printf(
		"Updated Weights → Risk: %.2f | Cost: %.2f | Penalty: %.2f\n",
		newWeights.RiskWeight,
		newWeights.CostWeight,
		newWeights.Penalty,
	)

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
