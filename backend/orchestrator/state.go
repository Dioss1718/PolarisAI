package orchestrator

import "sync"

var (
	stateMu     sync.RWMutex
	latestState *PipelineResult
)

func SetLatestState(state *PipelineResult) {
	stateMu.Lock()
	defer stateMu.Unlock()
	latestState = state
}

func GetLatestState() *PipelineResult {
	stateMu.RLock()
	defer stateMu.RUnlock()
	return latestState
}
