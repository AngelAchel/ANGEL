package sqlinject

import (
	"time"
)

type Sqlinject0109 struct{}

func NewSqlinject0109() *Sqlinject0109 {
	return &Sqlinject0109{}
}

func (e *Sqlinject0109) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0109) Name() string { return "Sqlinject0109" }
func (e *Sqlinject0109) Timestamp() time.Time { return time.Now() }
