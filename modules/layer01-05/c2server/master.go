package c2server

import (
	"time"
)

type Master struct{}

func NewMaster() *Master {
	return &Master{}
}

func (e *Master) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "master:done")
	return results, nil
}

func (e *Master) Name() string { return "Master" }
func (e *Master) Timestamp() time.Time { return time.Now() }
