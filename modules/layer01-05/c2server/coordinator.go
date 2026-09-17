package c2server

import (
	"time"
)

type Coordinator struct{}

func NewCoordinator() *Coordinator {
	return &Coordinator{}
}

func (e *Coordinator) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "coordinator:done")
	return results, nil
}

func (e *Coordinator) Name() string { return "Coordinator" }
func (e *Coordinator) Timestamp() time.Time { return time.Now() }
