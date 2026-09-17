package sqlinject

import (
	"time"
)

type Sqlinject0124 struct{}

func NewSqlinject0124() *Sqlinject0124 {
	return &Sqlinject0124{}
}

func (e *Sqlinject0124) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0124) Name() string { return "Sqlinject0124" }
func (e *Sqlinject0124) Timestamp() time.Time { return time.Now() }
