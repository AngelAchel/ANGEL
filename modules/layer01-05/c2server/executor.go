package c2server

import (
	"time"
)

type Executor struct{}

func NewExecutor() *Executor {
	return &Executor{}
}

func (e *Executor) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "executor:done")
	return results, nil
}

func (e *Executor) Name() string         { return "Executor" }
func (e *Executor) Timestamp() time.Time { return time.Now() }
