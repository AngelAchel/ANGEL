package sqlinject

import (
	"time"
)

type Sqlinject0150 struct{}

func NewSqlinject0150() *Sqlinject0150 {
	return &Sqlinject0150{}
}

func (e *Sqlinject0150) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0150) Name() string { return "Sqlinject0150" }
func (e *Sqlinject0150) Timestamp() time.Time { return time.Now() }
