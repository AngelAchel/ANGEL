package sqlinject

import (
	"time"
)

type Sqlinject0024 struct{}

func NewSqlinject0024() *Sqlinject0024 {
	return &Sqlinject0024{}
}

func (e *Sqlinject0024) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0024) Name() string { return "Sqlinject0024" }
func (e *Sqlinject0024) Timestamp() time.Time { return time.Now() }
