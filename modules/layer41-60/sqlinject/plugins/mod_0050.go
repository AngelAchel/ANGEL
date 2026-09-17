package sqlinject

import (
	"time"
)

type Sqlinject0050 struct{}

func NewSqlinject0050() *Sqlinject0050 {
	return &Sqlinject0050{}
}

func (e *Sqlinject0050) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0050) Name() string { return "Sqlinject0050" }
func (e *Sqlinject0050) Timestamp() time.Time { return time.Now() }
