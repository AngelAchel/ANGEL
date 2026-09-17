package c2server

import (
	"time"
)

type Job struct{}

func NewJob() *Job {
	return &Job{}
}

func (e *Job) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "job:done")
	return results, nil
}

func (e *Job) Name() string         { return "Job" }
func (e *Job) Timestamp() time.Time { return time.Now() }
