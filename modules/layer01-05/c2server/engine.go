package c2server

import (
	"time"
)

type Engine struct{}

func NewEngine() *Engine {
	return &Engine{}
}

func (e *Engine) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "engine:done")
	return results, nil
}

func (e *Engine) Name() string         { return "Engine" }
func (e *Engine) Timestamp() time.Time { return time.Now() }
