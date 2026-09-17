package sqlinject

import (
	"time"
)

type Sqlinject0093 struct{}

func NewSqlinject0093() *Sqlinject0093 {
	return &Sqlinject0093{}
}

func (e *Sqlinject0093) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0093) Name() string { return "Sqlinject0093" }
func (e *Sqlinject0093) Timestamp() time.Time { return time.Now() }
