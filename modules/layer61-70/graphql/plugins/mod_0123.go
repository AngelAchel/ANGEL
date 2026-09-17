package graphql

import (
    "time"
)

type graphql0123 struct{}

func Newgraphql0123() *graphql0123 {
    return &graphql0123{}
}

func (e *graphql0123) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "graphql:done")
    return results, nil
}

func (e *graphql0123) Name() string { return "graphql0123" }
func (e *graphql0123) Timestamp() time.Time { return time.Now() }
