package sqlinject

import (
	"time"
)

type Sqlinject0123 struct{}

func NewSqlinject0123() *Sqlinject0123 {
	return &Sqlinject0123{}
}

func (e *Sqlinject0123) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0123) Name() string { return "Sqlinject0123" }
func (e *Sqlinject0123) Timestamp() time.Time { return time.Now() }
