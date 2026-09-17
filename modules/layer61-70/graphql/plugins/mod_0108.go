package graphql

import (
    "time"
)

type graphql0108 struct{}

func Newgraphql0108() *graphql0108 {
    return &graphql0108{}
}

func (e *graphql0108) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "graphql:done")
    return results, nil
}

func (e *graphql0108) Name() string { return "graphql0108" }
func (e *graphql0108) Timestamp() time.Time { return time.Now() }
