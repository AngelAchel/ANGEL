package sqlinject

import (
	"time"
)

type Sqlinject0146 struct{}

func NewSqlinject0146() *Sqlinject0146 {
	return &Sqlinject0146{}
}

func (e *Sqlinject0146) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0146) Name() string { return "Sqlinject0146" }
func (e *Sqlinject0146) Timestamp() time.Time { return time.Now() }
