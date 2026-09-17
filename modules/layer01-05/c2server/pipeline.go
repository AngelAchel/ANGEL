package c2server

import (
	"time"
)

type Pipeline struct{}

func NewPipeline() *Pipeline {
	return &Pipeline{}
}

func (e *Pipeline) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "pipeline:done")
	return results, nil
}

func (e *Pipeline) Name() string         { return "Pipeline" }
func (e *Pipeline) Timestamp() time.Time { return time.Now() }
