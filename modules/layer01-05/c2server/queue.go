package c2server

import (
	"time"
)

type Queue struct{}

func NewQueue() *Queue {
	return &Queue{}
}

func (e *Queue) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "queue:done")
	return results, nil
}

func (e *Queue) Name() string         { return "Queue" }
func (e *Queue) Timestamp() time.Time { return time.Now() }
