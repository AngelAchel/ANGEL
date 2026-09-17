package c2server

import (
	"time"
)

type Processor struct{}

func NewProcessor() *Processor {
	return &Processor{}
}

func (e *Processor) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "processor:done")
	return results, nil
}

func (e *Processor) Name() string { return "Processor" }
func (e *Processor) Timestamp() time.Time { return time.Now() }
