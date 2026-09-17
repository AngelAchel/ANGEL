package c2server

import (
	"time"
)

// Orchestrator connects all layers
type Orchestrator struct{}

func NewOrchestrator() *Orchestrator {
	return &Orchestrator{}
}

func (o *Orchestrator) Orchestrate() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "orchestrator:all_layers_connected")
	return results, nil
}

func (o *Orchestrator) Name() string { return "Orchestrator" }
func (o *Orchestrator) Timestamp() time.Time { return time.Now() }
