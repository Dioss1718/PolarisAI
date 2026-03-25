
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
	policyvalidator "github.com/diya-suryawanshi/cloud/agents/policy-validator"
	securitysentinel "github.com/diya-suryawanshi/cloud/agents/security-sentinel"
	riskengine "github.com/diya-suryawanshi/cloud/agents/security-sentinel/risk-engine"
	"github.com/diya-suryawanshi/cloud/graph-engine/builder"
	"github.com/diya-suryawanshi/cloud/graph-engine/services"
	gitops "github.com/diya-suryawanshi/cloud/gitops"
)

func main() {
	fmt.Println("Starting Unified Cloud Graph Engine")
	startTime := time.Now()

	// ==============================
	// STEP 1: Fetch Data
	// ==============================
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

	// ==============================
	// STEP 2: Build Graph
	// ==============================
	g := builder.BuildGraph(parsed)

	fmt.Println("Graph successfully built")
	fmt.Printf("Nodes: %d\n", len(g.Nodes))
	fmt.Printf("Edges: %d\n", len(g.Edges))

	// ==============================
	// STEP 3: PARALLEL EXECUTION
	// ==============================
	var (
		attackPaths [][]string
		nodeRisks   map[string]float64
		signals     []costoptimizer.CostSignal
	)

	var wg sync.WaitGroup
	wg.Add(2)

	// 🔐 Security Sentinel
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

	// 💰 Cost Optimizer
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
	// STEP 4: CANDIDATE GENERATOR
	// ==============================
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

	// ==============================
	// STEP 5: ACTION GENERATOR
	// ==============================
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

	// ==============================
	// STEP 6: CONVERT → PARETO
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
	// STEP 7: PARETO OPTIMIZER
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
	// STEP 8: POLICY VALIDATOR (FIXED)
	// ==============================
	fmt.Println("\nRunning Policy Validator")

	policy := policyvalidator.Policy{
		MaxDowntime:        2.0,
		NoTerminateProd:    true,
		NoPublicDB:         true,
		EncryptionRequired: true,
	}

	for _, d := range decisions {

		node := g.Nodes[d.NodeID]

		// Compute centrality (same logic as elsewhere)
		centrality := float64(len(g.Adjacency[d.NodeID])) / 5.0

		// Get risk
		risk := nodeRisks[d.NodeID]

		input := policyvalidator.InputDecision{
			NodeID:        d.NodeID,
			Action:        d.Action,
			CostDelta:     0, // optional for now
			RiskReduction: 0, // optional for now
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

		// 🔥 FINAL DECISION LOGIC (IMPORTANT)
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

		fmt.Printf(
			"SLA: %.2f | Security: %.2f | Compliance: %.2f | Blast: %.2f\n",
			scores.SLA,
			scores.Security,
			scores.Compliance,
			scores.Blast,
		)
	}

// ==============================
// STEP 9: GITOPS INTEGRATION
// ==============================

validated := policyvalidator.RunPolicyValidator(g, []policyvalidator.InputDecision{}, nodeRisks)

var decisionsGitops []gitops.Decision

for _, v := range validated {
	if v.Status == "REJECTED" {
		continue
	}

	decisionsGitops = append(decisionsGitops, gitops.Decision{
		NodeID:      v.NodeID,
		Action:      v.Action,
		FinalAction: v.FinalAction,
		Score:       v.Score,
		Reason:      v.Reason,
	})
}

gitops.RunGitOps(g, decisionsGitops)

	// ==============================
	// FINAL SUMMARY
	// ==============================
	fmt.Println("\nExecution Summary")
	fmt.Printf("Total Time    : %v\n", time.Since(startTime))
	fmt.Printf("Attack Paths  : %d\n", len(attackPaths))
	fmt.Printf("Nodes Analyzed: %d\n", len(g.Nodes))

	fmt.Println("\nExecution Completed Successfully")
}