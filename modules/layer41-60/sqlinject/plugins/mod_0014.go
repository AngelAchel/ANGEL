package sqlinject

import (
	"time"
)

type Sqlinject0014 struct{}

func NewSqlinject0014() *Sqlinject0014 {
	return &Sqlinject0014{}
}

func (e *Sqlinject0014) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0014) Name() string { return "Sqlinject0014" }
func (e *Sqlinject0014) Timestamp() time.Time { return time.Now() }
