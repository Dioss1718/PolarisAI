package server

import (
	"sync"

	"github.com/diya-suryawanshi/cloud/backend/orchestrator"
)

type runtimeManager struct {
	mu           sync.Mutex
	lastScenario string
	lastSeed     int
	lastState    *orchestrator.PipelineResult
}

func newRuntimeManager() *runtimeManager {
	return &runtimeManager{}
}

func (rm *runtimeManager) setLatest(result *orchestrator.PipelineResult, scenario string, seed int) {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	rm.lastScenario = scenario
	rm.lastSeed = seed
	rm.lastState = result
}

func (rm *runtimeManager) latestFor(scenario string, seed int) *orchestrator.PipelineResult {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	if rm.lastState == nil {
		return nil
	}

	if rm.lastScenario != scenario || rm.lastSeed != seed {
		return nil
	}

	return rm.lastState
}
