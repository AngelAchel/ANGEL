package c2server

import (
	"time"
)

type Reducer struct{}

func NewReducer() *Reducer {
	return &Reducer{}
}

func (e *Reducer) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "reducer:done")
	return results, nil
}

func (e *Reducer) Name() string         { return "Reducer" }
func (e *Reducer) Timestamp() time.Time { return time.Now() }
