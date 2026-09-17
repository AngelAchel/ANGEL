package sqlinject

import (
	"time"
)

type Sqlinject0072 struct{}

func NewSqlinject0072() *Sqlinject0072 {
	return &Sqlinject0072{}
}

func (e *Sqlinject0072) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0072) Name() string { return "Sqlinject0072" }
func (e *Sqlinject0072) Timestamp() time.Time { return time.Now() }
