package graphql

import (
    "time"
)

type graphql0102 struct{}

func Newgraphql0102() *graphql0102 {
    return &graphql0102{}
}

func (e *graphql0102) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "graphql:done")
    return results, nil
}

func (e *graphql0102) Name() string { return "graphql0102" }
func (e *graphql0102) Timestamp() time.Time { return time.Now() }
