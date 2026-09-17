package storage

import (
	"time"
)

type Result struct{}

func NewResult() *Result {
	return &Result{}
}

func (r *Result) Get() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "result:done")
	return results, nil
}

func (r *Result) Name() string         { return "Result" }
func (r *Result) Timestamp() time.Time { return time.Now() }
