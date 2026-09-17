package c2server

import (
	"time"
)

type Analyzer struct{}

func NewAnalyzer() *Analyzer {
	return &Analyzer{}
}

func (e *Analyzer) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "analyzer:done")
	return results, nil
}

func (e *Analyzer) Name() string         { return "Analyzer" }
func (e *Analyzer) Timestamp() time.Time { return time.Now() }
