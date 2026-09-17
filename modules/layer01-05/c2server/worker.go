package c2server

import (
	"time"
)

type Worker struct{}

func NewWorker() *Worker {
	return &Worker{}
}

func (e *Worker) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "worker:done")
	return results, nil
}

func (e *Worker) Name() string { return "Worker" }
func (e *Worker) Timestamp() time.Time { return time.Now() }
