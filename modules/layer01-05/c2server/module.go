package c2server

import (
	"time"
)

type Module struct{}

func NewModule() *Module {
	return &Module{}
}

func (e *Module) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "module:done")
	return results, nil
}

func (e *Module) Name() string         { return "Module" }
func (e *Module) Timestamp() time.Time { return time.Now() }
