package c2server

import (
	"time"
)

type Graphql struct{}

func NewGraphql() *Graphql {
	return &Graphql{}
}

func (e *Graphql) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "graphql:done")
	return results, nil
}

func (e *Graphql) Name() string { return "Graphql" }
func (e *Graphql) Timestamp() time.Time { return time.Now() }
