package c2

import (
	"time"
)

type RouteDecide struct{}

func NewRouteDecide() *RouteDecide {
	return &RouteDecide{}
}

func (r *RouteDecide) Decide() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "route_decide:done")
	return results, nil
}

func (r *RouteDecide) Name() string         { return "RouteDecide" }
func (r *RouteDecide) Timestamp() time.Time { return time.Now() }
