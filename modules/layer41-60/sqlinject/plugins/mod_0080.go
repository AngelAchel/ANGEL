package sqlinject

import (
	"time"
)

type Sqlinject0080 struct{}

func NewSqlinject0080() *Sqlinject0080 {
	return &Sqlinject0080{}
}

func (e *Sqlinject0080) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0080) Name() string { return "Sqlinject0080" }
func (e *Sqlinject0080) Timestamp() time.Time { return time.Now() }
