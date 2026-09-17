package c2server

import (
	"time"
)

type Runner struct{}

func NewRunner() *Runner {
	return &Runner{}
}

func (e *Runner) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "runner:done")
	return results, nil
}

func (e *Runner) Name() string { return "Runner" }
func (e *Runner) Timestamp() time.Time { return time.Now() }
