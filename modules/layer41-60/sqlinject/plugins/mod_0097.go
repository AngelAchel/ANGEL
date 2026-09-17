package sqlinject

import (
	"time"
)

type Sqlinject0097 struct{}

func NewSqlinject0097() *Sqlinject0097 {
	return &Sqlinject0097{}
}

func (e *Sqlinject0097) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0097) Name() string { return "Sqlinject0097" }
func (e *Sqlinject0097) Timestamp() time.Time { return time.Now() }
