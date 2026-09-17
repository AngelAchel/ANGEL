package c2server

import (
	"time"
)

type Task struct{}

func NewTask() *Task {
	return &Task{}
}

func (e *Task) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "task:done")
	return results, nil
}

func (e *Task) Name() string { return "Task" }
func (e *Task) Timestamp() time.Time { return time.Now() }
