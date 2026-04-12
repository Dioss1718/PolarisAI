package server

import (
	"fmt"
	"math"
	"strings"

	"github.com/diya-suryawanshi/cloud/backend/orchestrator"
)

type compareRequest struct {
	BaseScenario      string                 `json:"baseScenario"`
	BaseSeed          int                    `json:"baseSeed"`
	CandidateScenario string                 `json:"candidateScenario"`
	CandidateSeed     int                    `json:"candidateSeed"`
	ManualData        map[string]interface{} `json:"manualData,omitempty"`
}

type compareDelta struct {
	CurrentTotalCostDelta    float64 `json:"currentTotalCostDelta"`
	Forecast30TotalDelta     float64 `json:"forecast30TotalDelta"`
	Forecast90TotalDelta     float64 `json:"forecast90TotalDelta"`
	AttackPathCountDelta     int     `json:"attackPathCountDelta"`
	PublicExposureCountDelta int     `json:"publicExposureCountDelta"`
	HighRiskCountDelta       int     `json:"highRiskCountDelta"`
	BillShockCountDelta      int     `json:"billShockCountDelta"`
	AverageRiskDelta         float64 `json:"averageRiskDelta"`
	ComplianceScoreDelta     float64 `json:"complianceScoreDelta"`
	CostRiskScoreDelta       float64 `json:"costRiskScoreDelta"`
}

type compareResponse struct {
	Baseline  *orchestrator.PipelineResult `json:"baseline"`
	Candidate *orchestrator.PipelineResult `json:"candidate"`
	Delta     compareDelta                 `json:"delta"`
	Mode      string                       `json:"mode"`
}

func runWhatIfCompare(req compareRequest) (*compareResponse, error) {
	baseScenario := strings.TrimSpace(req.BaseScenario)
	if baseScenario == "" {
		baseScenario = "FULL_CHAOS"
	}

	baseSeed := req.BaseSeed
	if baseSeed == 0 {
		baseSeed = 42
	}

	candidateScenario := strings.TrimSpace(req.CandidateScenario)
	if candidateScenario == "" {
		candidateScenario = baseScenario
	}

	candidateSeed := req.CandidateSeed
	if candidateSeed == 0 {
		candidateSeed = baseSeed
	}

	baseline, err := orchestrator.Run(orchestrator.RunRequest{
		Scenario:           baseScenario,
		Seed:               baseSeed,
		SkipExplainability: true,
	})
	if err != nil {
		return nil, fmt.Errorf("baseline run failed: %w", err)
	}

	candidateReq := orchestrator.RunRequest{
		Scenario:           candidateScenario,
		Seed:               candidateSeed,
		SkipExplainability: true,
	}

	mode := "scenario_vs_scenario"

	if req.ManualData != nil {
		candidateReq = orchestrator.RunRequest{
			Scenario:           "MANUAL",
			Seed:               candidateSeed,
			ManualData:         req.ManualData,
			SkipExplainability: true,
		}
		mode = "scenario_vs_manual"
	}

	candidate, err := orchestrator.Run(candidateReq)
	if err != nil {
		return nil, fmt.Errorf("candidate run failed: %w", err)
	}

	return &compareResponse{
		Baseline:  baseline,
		Candidate: candidate,
		Delta:     buildCompareDelta(baseline, candidate),
		Mode:      mode,
	}, nil
}

func buildCompareDelta(base, candidate *orchestrator.PipelineResult) compareDelta {
	if base == nil || candidate == nil {
		return compareDelta{}
	}

	return compareDelta{
		CurrentTotalCostDelta:    round2(candidate.Summary.CurrentTotalCost - base.Summary.CurrentTotalCost),
		Forecast30TotalDelta:     round2(candidate.Summary.Forecast30Total - base.Summary.Forecast30Total),
		Forecast90TotalDelta:     round2(candidate.Summary.Forecast90Total - base.Summary.Forecast90Total),
		AttackPathCountDelta:     candidate.Summary.AttackPathCount - base.Summary.AttackPathCount,
		PublicExposureCountDelta: candidate.Summary.PublicExposureCount - base.Summary.PublicExposureCount,
		HighRiskCountDelta:       candidate.Summary.HighRiskCount - base.Summary.HighRiskCount,
		BillShockCountDelta:      candidate.Summary.BillShockCount - base.Summary.BillShockCount,
		AverageRiskDelta:         round2(candidate.Summary.AverageRisk - base.Summary.AverageRisk),
		ComplianceScoreDelta:     round2(candidate.Summary.ComplianceScore - base.Summary.ComplianceScore),
		CostRiskScoreDelta:       round2(candidate.Summary.CostRiskScore - base.Summary.CostRiskScore),
	}
}

func round2(v float64) float64 {
	return math.Round(v*100) / 100
}
