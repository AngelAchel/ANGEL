package c2server

import (
	"time"
)

type Operator struct{}

func NewOperator() *Operator {
	return &Operator{}
}

func (e *Operator) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "operator:done")
	return results, nil
}

func (e *Operator) Name() string { return "Operator" }
func (e *Operator) Timestamp() time.Time { return time.Now() }
