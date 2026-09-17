package graphql

import (
    "time"
)

type graphql0100 struct{}

func Newgraphql0100() *graphql0100 {
    return &graphql0100{}
}

func (e *graphql0100) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "graphql:done")
    return results, nil
}

func (e *graphql0100) Name() string { return "graphql0100" }
func (e *graphql0100) Timestamp() time.Time { return time.Now() }
