package sqlinject

import (
	"time"
)

type Sqlinject0112 struct{}

func NewSqlinject0112() *Sqlinject0112 {
	return &Sqlinject0112{}
}

func (e *Sqlinject0112) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0112) Name() string { return "Sqlinject0112" }
func (e *Sqlinject0112) Timestamp() time.Time { return time.Now() }
