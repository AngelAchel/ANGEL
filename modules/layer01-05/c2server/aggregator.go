package c2server

import (
	"time"
)

type Aggregator struct{}

func NewAggregator() *Aggregator {
	return &Aggregator{}
}

func (e *Aggregator) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "aggregator:done")
	return results, nil
}

func (e *Aggregator) Name() string         { return "Aggregator" }
func (e *Aggregator) Timestamp() time.Time { return time.Now() }
