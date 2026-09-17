package c2server

import (
	"time"
)

type Collector struct{}

func NewCollector() *Collector {
	return &Collector{}
}

func (e *Collector) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "collector:done")
	return results, nil
}

func (e *Collector) Name() string { return "Collector" }
func (e *Collector) Timestamp() time.Time { return time.Now() }
