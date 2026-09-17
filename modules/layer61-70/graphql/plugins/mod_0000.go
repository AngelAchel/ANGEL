package graphql

import (
    "time"
)

type graphql0000 struct{}

func Newgraphql0000() *graphql0000 {
    return &graphql0000{}
}

func (e *graphql0000) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "graphql:done")
    return results, nil
}

func (e *graphql0000) Name() string { return "graphql0000" }
func (e *graphql0000) Timestamp() time.Time { return time.Now() }
