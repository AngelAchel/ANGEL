package sqlinject

import (
	"time"
)

type Sqlinject0137 struct{}

func NewSqlinject0137() *Sqlinject0137 {
	return &Sqlinject0137{}
}

func (e *Sqlinject0137) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0137) Name() string { return "Sqlinject0137" }
func (e *Sqlinject0137) Timestamp() time.Time { return time.Now() }
