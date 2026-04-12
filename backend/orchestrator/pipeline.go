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

	"github.com/diya-suryawanshi/cloud/carbon"
	forecast "github.com/diya-suryawanshi/cloud/forecast"
	gitops "github.com/diya-suryawanshi/cloud/gitops"
	"github.com/diya-suryawanshi/cloud/graph-engine/builder"
	"github.com/diya-suryawanshi/cloud/graph-engine/services"
)

func Run(req RunRequest) (*PipelineResult, error) {
	startTime := time.Now()

	scenario := req.Scenario
	if scenario == "" {
		scenario = "FULL_CHAOS"
	}

	seed := req.Seed
	if seed == 0 {
		seed = 42
	}

	var parsed map[string]interface{}

	if req.ManualData != nil {
		parsed = req.ManualData
	} else {
		data, err := services.FetchSimulationData(scenario, seed)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch simulation data: %w", err)
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal simulation data: %w", err)
		}

		if err := json.Unmarshal(jsonData, &parsed); err != nil {
			return nil, fmt.Errorf("failed to unmarshal simulation data: %w", err)
		}
	}

	g, err := builder.BuildGraph(parsed)
	if err != nil {
		return nil, fmt.Errorf("invalid graph input: %w", err)
	}

	stages := []PipelineStageDTO{
		{Name: "Graph Build", Status: "complete"},
		{Name: "Risk Modeling", Status: "running"},
	}

	var (
		nodeRisks   map[string]float64
		attackPaths [][]string
		signals     []costoptimizer.CostSignal
	)

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		nodeRisks = riskengine.ComputeNodeRisk(g)
	}()

	go func() {
		defer wg.Done()
		attackPaths = securitysentinel.FindAttackPaths(g)
	}()

	go func() {
		defer wg.Done()
		signals = costoptimizer.Run(g)
	}()

	wg.Wait()

	attackMetrics := ComputeAttackMetrics(attackPaths)

	stages = append(stages,
		PipelineStageDTO{Name: "Security Sentinel", Status: "complete"},
		PipelineStageDTO{Name: "Cost Optimizer", Status: "complete"},
	)

	candidates := candidategenerator.GenerateCandidates(g, signals, nodeRisks)
	stages = append(stages, PipelineStageDTO{Name: "Candidate Generator", Status: "complete"})

	actions := actiongenerator.GenerateActions(g, candidates)
	stages = append(stages, PipelineStageDTO{Name: "Action Generator", Status: "complete"})

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
	stages = append(stages, PipelineStageDTO{Name: "Pareto Optimizer", Status: "complete"})

	blastMap := computeBlastRadius(g)
	policyInputs := make([]policyvalidator.InputDecision, 0, len(decisions))
	for _, d := range decisions {
		srcAction := actionLookup[d.NodeID+"|"+d.Action]

		policyInputs = append(policyInputs, policyvalidator.InputDecision{
			NodeID:        d.NodeID,
			Action:        d.Action,
			CostDelta:     srcAction.CostDelta,
			RiskReduction: srcAction.RiskReduction,
		})
	}

	validatedDecisions := policyvalidator.RunPolicyValidator(g, policyInputs, nodeRisks)

	validatedLookup := make(map[string]policyvalidator.ValidatedDecision, len(validatedDecisions))
	for _, vd := range validatedDecisions {
		key := vd.NodeID + "|" + vd.Action
		validatedLookup[key] = vd
	}

	var recommendations []RecommendationDTO
	var decisionsGitops []gitops.Decision
	var approvedActions []actiongenerator.Action
	var explanations []ExplanationDTO
	var forecasts []ForecastDTO

	for _, d := range decisions {
		node := g.Nodes[d.NodeID]
		risk := nodeRisks[d.NodeID]

		validation, ok := validatedLookup[d.NodeID+"|"+d.Action]
		if !ok {
			validation = policyvalidator.ValidatedDecision{
				Status:      "REJECTED",
				FinalAction: d.Action,
				Score:       0,
				Reason:      "Policy validation missing for action",
				Scores: policyvalidator.ValidationScores{
					SLA:        0,
					Security:   0,
					Compliance: 0,
					Blast:      0,
				},
			}
		}

		scores := validation.Scores
		finalScore := validation.Score
		status := validation.Status
		finalAction := validation.FinalAction

		srcAction := actionLookup[d.NodeID+"|"+d.Action]
		confidence := deriveConfidence(srcAction.Score, finalScore, node.Environment)
		safetyLevel := deriveSafetyLevel(node.Environment, node.Exposure, finalScore)

		rec := RecommendationDTO{
			NodeID:         d.NodeID,
			Action:         d.Action,
			FinalAction:    finalAction,
			Status:         status,
			Score:          round2(finalScore),
			Reason:         d.Reason + " | " + validation.Reason,
			Risk:           round2(risk),
			Cloud:          node.Cloud,
			Type:           node.Type,
			Environment:    node.Environment,
			CostDelta:      round2(srcAction.CostDelta),
			RiskReduction:  round2(srcAction.RiskReduction),
			Confidence:     confidence,
			SafetyLevel:    safetyLevel,
			BlastRadius:    blastMap[d.NodeID],
			ComplianceGain: round2(scores.Compliance),
			GitOpsPath:     fmt.Sprintf("GitOps PR -> branch polaris/%s/%s", d.NodeID, finalAction),
			RollbackPath:   fmt.Sprintf("Rollback via revert of branch polaris/%s/%s", d.NodeID, finalAction),
		}
		recommendations = append(recommendations, rec)

		if status == "REJECTED" {
			continue
		}

		decisionsGitops = append(decisionsGitops, gitops.Decision{
			NodeID:      d.NodeID,
			Action:      d.Action,
			FinalAction: finalAction,
			Score:       finalScore,
			Reason:      validation.Reason,
		})

		if src, ok := actionLookup[d.NodeID+"|"+d.Action]; ok {
			approvedActions = append(approvedActions, src)
		}

		explainReq := aiexplain.AIRequest{
			NodeID:        d.NodeID,
			Action:        finalAction,
			Env:           node.Environment,
			NodeType:      node.Type,
			Cost:          srcAction.CostDelta,
			RiskReduction: srcAction.RiskReduction,
			SLA:           scores.SLA,
			Security:      scores.Security,
			Compliance:    scores.Compliance,
			Blast:         scores.Blast,
		}

		explanationResp, err := aiexplain.GetExplanation(explainReq)
		if err == nil {
			explanations = append(explanations, ExplanationDTO{
				NodeID:      d.NodeID,
				Action:      finalAction,
				Explanation: explanationResp.Explanation,
				Grounded:    explanationResp.Grounded,
				Sources:     explanationResp.Sources,
			})
		}

		if f, err := forecast.Get(d.NodeID); err == nil {
			forecasts = append(forecasts, ForecastDTO{
				NodeID:      f.NodeID,
				CurrentCost: f.CurrentCost,
				Forecast30:  f.Forecast30,
				Forecast90:  f.Forecast90,
				BillShock:   f.BillShock,
				ShockReason: f.ShockReason,
			})
		}
	}

	stages = append(stages,
		PipelineStageDTO{Name: "Policy Validator", Status: "complete"},
		PipelineStageDTO{Name: "Forecast", Status: "complete"},
		PipelineStageDTO{Name: "Explainability", Status: "complete"},
	)

	gitopsStatus := GitOpsDTO{
		Status:  "pending_approval",
		Message: fmt.Sprintf("%d validated remediation(s) are awaiting human approval before PR creation", len(decisionsGitops)),
		PRs:     []GitOpsPRDTO{},
	}

	stages = append(stages, PipelineStageDTO{Name: "GitOps", Status: "pending_approval"})
	records := feedback.Load()
	for _, a := range approvedActions {
		record := feedback.CreateRecord(
			a.NodeID,
			a.ActionType+"_"+a.Variant,
			a.CostDelta,
			a.RiskReduction,
			a.Score,
		)
		records = append(records, record)
	}
	feedback.Save(records)

	feedbackSummary := feedback.Summarize(records)
	newWeights := feedback.UpdateWeights(feedbackSummary)
	stages = append(stages, PipelineStageDTO{Name: "Feedback Learning", Status: "complete"})

	finalActions := make(map[string]RecommendationDTO)
	for _, r := range recommendations {
		finalActions[r.NodeID] = r
	}

	var nodes []GraphNodeDTO
	for _, n := range g.Nodes {
		nodeDTO := GraphNodeDTO{
			ID:          n.ID,
			Label:       n.Name,
			Type:        n.Type,
			Cloud:       n.Cloud,
			Region:      n.Region,
			Environment: n.Environment,
			Cost:        n.Cost,
			Utilization: n.Utilization,
			Exposure:    n.Exposure,
			Criticality: n.Criticality,
			Compliance:  n.Compliance,
			Risk:        nodeRisks[n.ID],
		}

		if rec, ok := finalActions[n.ID]; ok {
			nodeDTO.FinalAction = rec.FinalAction
			nodeDTO.Status = rec.Status
		}

		nodes = append(nodes, nodeDTO)
	}

	var edges []GraphEdgeDTO
	for _, e := range g.Edges {
		edges = append(edges, GraphEdgeDTO{
			From:   e.From,
			To:     e.To,
			Type:   e.Type,
			Weight: e.Weight,
		})
	}

	currentCarbon := carbon.Run(carbon.FromGraphNodes(g.Nodes))

	projectedGraph := gitops.GenerateFullProposedGraph(g, decisionsGitops)
	projectedAttackPaths := securitysentinel.FindAttackPaths(projectedGraph)
	projectedAttackMetrics := ComputeAttackMetrics(projectedAttackPaths)
	projectedRisks := riskengine.ComputeNodeRisk(projectedGraph)
	projectedCarbon := carbon.Run(carbon.FromGraphNodes(projectedGraph.Nodes))

	summary := BuildCurrentSummary(
		g,
		attackMetrics,
		recommendations,
		forecasts,
		nodeRisks,
		currentCarbon,
	)

	projectedSummary := BuildProjectedSummary(
		projectedGraph,
		projectedAttackMetrics,
		projectedRisks,
		projectedCarbon,
		currentCarbon.Total,
		summary.AverageRisk,
	)

	nodeIntel := buildNodeIntel(g, nodeRisks, attackPaths, forecasts)
	edgeInfluence := buildEdgeInfluence(g, nodeRisks)
	negotiations := buildNegotiationTraces(actions, recommendations)
	alerts := buildStructuredAlerts(summary, projectedSummary, recommendations, nodeIntel)

	result := &PipelineResult{
		Scenario:        scenario,
		Seed:            seed,
		Nodes:           nodes,
		Edges:           edges,
		Recommendations: recommendations,
		Explanations:    explanations,
		Forecasts:       forecasts,
		Feedback: FeedbackDTO{
			AvgReward:  feedbackSummary.AvgReward,
			Count:      feedbackSummary.Count,
			RiskWeight: newWeights.RiskWeight,
			CostWeight: newWeights.CostWeight,
			Penalty:    newWeights.Penalty,
		},
		GitOps:      gitopsStatus,
		AttackPaths: attackPaths,
		AttackMetrics: AttackMetricsDTO{
			PathCount:      attackMetrics.PathCount,
			AvgPathLength:  attackMetrics.AvgPathLength,
			ReachableNodes: attackMetrics.ReachableNodes,
		},
		Summary:          summary,
		ProjectedSummary: projectedSummary,
		Stages:           stages,
		Alerts:           alerts,
		NodeIntel:        nodeIntel,
		EdgeInfluence:    edgeInfluence,
		Negotiations:     negotiations,
		GeneratedAt:      startTime.UTC().Format(time.RFC3339),
	}

	SetLatestState(result)
	return result, nil
}
