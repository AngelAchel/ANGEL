package sqlinject

import (
	"time"
)

type Sqlinject0033 struct{}

func NewSqlinject0033() *Sqlinject0033 {
	return &Sqlinject0033{}
}

func (e *Sqlinject0033) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0033) Name() string { return "Sqlinject0033" }
func (e *Sqlinject0033) Timestamp() time.Time { return time.Now() }
