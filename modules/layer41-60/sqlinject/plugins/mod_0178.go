package sqlinject

import (
	"time"
)

type Sqlinject0178 struct{}

func NewSqlinject0178() *Sqlinject0178 {
	return &Sqlinject0178{}
}

func (e *Sqlinject0178) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0178) Name() string { return "Sqlinject0178" }
func (e *Sqlinject0178) Timestamp() time.Time { return time.Now() }
