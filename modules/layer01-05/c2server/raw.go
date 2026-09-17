package c2server

import (
	"time"
)

type Raw struct{}

func NewRaw() *Raw {
	return &Raw{}
}

func (e *Raw) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "raw:done")
	return results, nil
}

func (e *Raw) Name() string { return "Raw" }
func (e *Raw) Timestamp() time.Time { return time.Now() }
