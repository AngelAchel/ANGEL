package server
//nolint:staticcheck

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"sync"
	"time"
)

type ResultHandler struct {
	mu       sync.RWMutex
	results  map[string]*TaskResult
	handlers map[string]ResultCallback
	running  bool
}

type TaskResult struct {
	ID        string
	TaskID    string
	AgentID   string
	Success   bool
	Output    string
	Error     string
	Timestamp time.Time
	Duration  time.Duration
	Data      string
	Type      string
}

type ResultCallback func(result *TaskResult)

type ResultHandlerConfig struct {
	MaxResults int
	BufferSize int
}

func NewResultHandler(config ResultHandlerConfig) *ResultHandler {
	return &ResultHandler{
		results:  make(map[string]*TaskResult),
		handlers: make(map[string]ResultCallback),
	}
}

func (rh *ResultHandler) HandleResult(result *TaskResult) {
	rh.mu.Lock()
	defer rh.mu.Unlock()

	if result.ID == "" {
		result.ID = generateResultID()
	}
	result.Timestamp = time.Now()

	rh.results[result.ID] = result

	handlersCopy := make(map[string]ResultCallback, len(rh.handlers))
	for k, v := range rh.handlers {
		handlersCopy[k] = v
	}

	for _, handler := range handlersCopy {
		r := &TaskResult{
			ID:        result.ID,
			TaskID:    result.TaskID,
			AgentID:   result.AgentID,
			Success:   result.Success,
			Output:    result.Output,
			Error:     result.Error,
			Timestamp: result.Timestamp,
			Duration:  result.Duration,
			Data:      result.Data,
			Type:      result.Type,
		}
		handler(r)
	}
}

func (rh *ResultHandler) GetResult(resultID string) *TaskResult {
	rh.mu.RLock()
	defer rh.mu.RUnlock()
	return rh.results[resultID]
}

func (rh *ResultHandler) GetResultsByTask(taskID string) []*TaskResult {
	rh.mu.RLock()
	defer rh.mu.RUnlock()

	results := make([]*TaskResult, 0)
	for _, result := range rh.results {
		if result.TaskID == taskID {
			results = append(results, result)
		}
	}
	return results
}

func (rh *ResultHandler) GetResultsByAgent(agentID string) []*TaskResult {
	rh.mu.RLock()
	defer rh.mu.RUnlock()

	results := make([]*TaskResult, 0)
	for _, result := range rh.results {
		if result.AgentID == agentID {
			results = append(results, result)
		}
	}
	return results
}

func (rh *ResultHandler) GetSuccessfulResults() []*TaskResult {
	rh.mu.RLock()
	defer rh.mu.RUnlock()

	results := make([]*TaskResult, 0)
	for _, result := range rh.results {
		if result.Success {
			results = append(results, result)
		}
	}
	return results
}

func (rh *ResultHandler) GetFailedResults() []*TaskResult {
	rh.mu.RLock()
	defer rh.mu.RUnlock()

	results := make([]*TaskResult, 0)
	for _, result := range rh.results {
		if !result.Success {
			results = append(results, result)
		}
	}
	return results
}

func (rh *ResultHandler) RegisterHandler(name string, callback ResultCallback) {
	rh.mu.Lock()
	defer rh.mu.Unlock()
	rh.handlers[name] = callback
}

func (rh *ResultHandler) UnregisterHandler(name string) {
	rh.mu.Lock()
	defer rh.mu.Unlock()
	delete(rh.handlers, name)
}

func (rh *ResultHandler) GetResultCount() int {
	rh.mu.RLock()
	defer rh.mu.RUnlock()
	return len(rh.results)
}

func (rh *ResultHandler) GetSuccessfulCount() int {
	rh.mu.RLock()
	defer rh.mu.RUnlock()

	count := 0
	for _, result := range rh.results {
		if result.Success {
			count++
		}
	}
	return count
}

func (rh *ResultHandler) GetFailedCount() int {
	rh.mu.RLock()
	defer rh.mu.RUnlock()

	count := 0
	for _, result := range rh.results {
		if !result.Success {
			count++
		}
	}
	return count
}

func (rh *ResultHandler) ClearResults() {
	rh.mu.Lock()
	defer rh.mu.Unlock()
	rh.results = make(map[string]*TaskResult)
}

func (rh *ResultHandler) IsRunning() bool {
	rh.mu.RLock()
	defer rh.mu.RUnlock()
	return rh.running
}

func generateResultID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}  //nolint:staticcheck
  //nolint:staticcheck
func formatResult(result *TaskResult) string {
	status := "✅ Success"
	if !result.Success {
		status = "❌ Failed"
	}

	return fmt.Sprintf("[%s] Task %s: %s (Duration: %v)",
		status, result.TaskID, result.Output, result.Duration)
}
