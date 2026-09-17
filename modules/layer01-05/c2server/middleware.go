package c2server

import (
	"time"
)

type Middleware struct{}

func NewMiddleware() *Middleware {
	return &Middleware{}
}

func (e *Middleware) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "middleware:done")
	return results, nil
}

func (e *Middleware) Name() string         { return "Middleware" }
func (e *Middleware) Timestamp() time.Time { return time.Now() }
