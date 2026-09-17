package c2server

import (
	"time"
)

type Dispatcher struct{}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{}
}

func (e *Dispatcher) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "dispatcher:done")
	return results, nil
}

func (e *Dispatcher) Name() string         { return "Dispatcher" }
func (e *Dispatcher) Timestamp() time.Time { return time.Now() }
