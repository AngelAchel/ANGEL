package c2server

import (
	"time"
)

type Injector struct{}

func NewInjector() *Injector {
	return &Injector{}
}

func (e *Injector) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "injector:done")
	return results, nil
}

func (e *Injector) Name() string         { return "Injector" }
func (e *Injector) Timestamp() time.Time { return time.Now() }
