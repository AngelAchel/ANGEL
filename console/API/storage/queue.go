package storage

import (
	"time"
)

type Queue struct{}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "queue:done")
	return results, nil
}

func (q *Queue) Name() string { return "Queue" }
func (q *Queue) Timestamp() time.Time { return time.Now() }
