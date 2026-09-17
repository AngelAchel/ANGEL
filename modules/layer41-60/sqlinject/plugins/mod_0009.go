package sqlinject

import (
	"time"
)

type Sqlinject0009 struct{}

func NewSqlinject0009() *Sqlinject0009 {
	return &Sqlinject0009{}
}

func (e *Sqlinject0009) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0009) Name() string { return "Sqlinject0009" }
func (e *Sqlinject0009) Timestamp() time.Time { return time.Now() }
