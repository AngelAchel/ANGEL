package c2server

import (
	"time"
)

type Filter struct{}

func NewFilter() *Filter {
	return &Filter{}
}

func (e *Filter) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "filter:done")
	return results, nil
}

func (e *Filter) Name() string         { return "Filter" }
func (e *Filter) Timestamp() time.Time { return time.Now() }
