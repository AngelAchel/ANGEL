package c2server

import (
	"time"
)

type Commander struct{}

func NewCommander() *Commander {
	return &Commander{}
}

func (e *Commander) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "commander:done")
	return results, nil
}

func (e *Commander) Name() string         { return "Commander" }
func (e *Commander) Timestamp() time.Time { return time.Now() }
