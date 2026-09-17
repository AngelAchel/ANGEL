package sqlinject

import (
	"time"
)

type Sqlinject0118 struct{}

func NewSqlinject0118() *Sqlinject0118 {
	return &Sqlinject0118{}
}

func (e *Sqlinject0118) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0118) Name() string { return "Sqlinject0118" }
func (e *Sqlinject0118) Timestamp() time.Time { return time.Now() }
