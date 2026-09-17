package sqlinject

import (
	"time"
)

type Sqlinject0051 struct{}

func NewSqlinject0051() *Sqlinject0051 {
	return &Sqlinject0051{}
}

func (e *Sqlinject0051) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0051) Name() string { return "Sqlinject0051" }
func (e *Sqlinject0051) Timestamp() time.Time { return time.Now() }
