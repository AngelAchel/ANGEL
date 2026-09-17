package c2server

import (
	"time"
)

type Batch struct{}

func NewBatch() *Batch {
	return &Batch{}
}

func (e *Batch) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "batch:done")
	return results, nil
}

func (e *Batch) Name() string         { return "Batch" }
func (e *Batch) Timestamp() time.Time { return time.Now() }
