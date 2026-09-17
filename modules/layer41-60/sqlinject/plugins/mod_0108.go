package sqlinject

import (
	"time"
)

type Sqlinject0108 struct{}

func NewSqlinject0108() *Sqlinject0108 {
	return &Sqlinject0108{}
}

func (e *Sqlinject0108) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0108) Name() string { return "Sqlinject0108" }
func (e *Sqlinject0108) Timestamp() time.Time { return time.Now() }
