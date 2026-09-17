package sqlinject

import (
	"time"
)

type Sqlinject0193 struct{}

func NewSqlinject0193() *Sqlinject0193 {
	return &Sqlinject0193{}
}

func (e *Sqlinject0193) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0193) Name() string { return "Sqlinject0193" }
func (e *Sqlinject0193) Timestamp() time.Time { return time.Now() }
