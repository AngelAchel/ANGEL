package graphql

import (
    "time"
)

type graphql0013 struct{}

func Newgraphql0013() *graphql0013 {
    return &graphql0013{}
}

func (e *graphql0013) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "graphql:done")
    return results, nil
}

func (e *graphql0013) Name() string { return "graphql0013" }
func (e *graphql0013) Timestamp() time.Time { return time.Now() }
