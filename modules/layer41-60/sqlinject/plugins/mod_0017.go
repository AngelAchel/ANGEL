package sqlinject

import (
	"time"
)

type Sqlinject0017 struct{}

func NewSqlinject0017() *Sqlinject0017 {
	return &Sqlinject0017{}
}

func (e *Sqlinject0017) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0017) Name() string { return "Sqlinject0017" }
func (e *Sqlinject0017) Timestamp() time.Time { return time.Now() }
