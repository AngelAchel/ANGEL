package c2server

import (
	"time"
)

type Scheduler struct{}

func NewScheduler() *Scheduler {
	return &Scheduler{}
}

func (e *Scheduler) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "scheduler:done")
	return results, nil
}

func (e *Scheduler) Name() string { return "Scheduler" }
func (e *Scheduler) Timestamp() time.Time { return time.Now() }
