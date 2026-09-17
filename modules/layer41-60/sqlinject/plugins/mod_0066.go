package sqlinject

import (
	"time"
)

type Sqlinject0066 struct{}

func NewSqlinject0066() *Sqlinject0066 {
	return &Sqlinject0066{}
}

func (e *Sqlinject0066) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0066) Name() string { return "Sqlinject0066" }
func (e *Sqlinject0066) Timestamp() time.Time { return time.Now() }
