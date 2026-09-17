package ai

import (
	"time"
)

type GraphQLParam struct{}

func NewGraphQLParam() *GraphQLParam {
	return &GraphQLParam{}
}

func (g *GraphQLParam) Query() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "graphql_param:done")
	return results, nil
}

func (g *GraphQLParam) Name() string { return "GraphQLParam" }
func (g *GraphQLParam) Timestamp() time.Time { return time.Now() }
