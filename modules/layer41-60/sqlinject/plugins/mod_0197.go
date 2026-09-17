package sqlinject

import (
	"time"
)

type Sqlinject0197 struct{}

func NewSqlinject0197() *Sqlinject0197 {
	return &Sqlinject0197{}
}

func (e *Sqlinject0197) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0197) Name() string { return "Sqlinject0197" }
func (e *Sqlinject0197) Timestamp() time.Time { return time.Now() }
