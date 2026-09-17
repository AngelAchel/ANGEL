package c2server

import (
	"time"
)

type Rest struct{}

func NewRest() *Rest {
	return &Rest{}
}

func (e *Rest) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "rest:done")
	return results, nil
}

func (e *Rest) Name() string { return "Rest" }
func (e *Rest) Timestamp() time.Time { return time.Now() }
