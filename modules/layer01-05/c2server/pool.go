package c2server

import (
	"time"
)

type Pool struct{}

func NewPool() *Pool {
	return &Pool{}
}

func (e *Pool) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "pool:done")
	return results, nil
}

func (e *Pool) Name() string         { return "Pool" }
func (e *Pool) Timestamp() time.Time { return time.Now() }
