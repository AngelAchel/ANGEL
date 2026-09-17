package graphql

import (
    "time"
)

type graphql0120 struct{}

func Newgraphql0120() *graphql0120 {
    return &graphql0120{}
}

func (e *graphql0120) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "graphql:done")
    return results, nil
}

func (e *graphql0120) Name() string { return "graphql0120" }
func (e *graphql0120) Timestamp() time.Time { return time.Now() }
