package c2server

import (
	"time"
)

type Workflow struct{}

func NewWorkflow() *Workflow {
	return &Workflow{}
}

func (e *Workflow) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "workflow:done")
	return results, nil
}

func (e *Workflow) Name() string { return "Workflow" }
func (e *Workflow) Timestamp() time.Time { return time.Now() }
