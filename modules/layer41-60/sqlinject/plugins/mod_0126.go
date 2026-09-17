package sqlinject

import (
	"time"
)

type Sqlinject0126 struct{}

func NewSqlinject0126() *Sqlinject0126 {
	return &Sqlinject0126{}
}

func (e *Sqlinject0126) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0126) Name() string { return "Sqlinject0126" }
func (e *Sqlinject0126) Timestamp() time.Time { return time.Now() }
