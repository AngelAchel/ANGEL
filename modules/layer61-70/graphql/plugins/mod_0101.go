package graphql

import (
    "time"
)

type graphql0101 struct{}

func Newgraphql0101() *graphql0101 {
    return &graphql0101{}
}

func (e *graphql0101) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "graphql:done")
    return results, nil
}

func (e *graphql0101) Name() string { return "graphql0101" }
func (e *graphql0101) Timestamp() time.Time { return time.Now() }
