package c2server

import (
	"time"
)

type Walker struct{}

func NewWalker() *Walker {
	return &Walker{}
}

func (e *Walker) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "walker:done")
	return results, nil
}

func (e *Walker) Name() string { return "Walker" }
func (e *Walker) Timestamp() time.Time { return time.Now() }
