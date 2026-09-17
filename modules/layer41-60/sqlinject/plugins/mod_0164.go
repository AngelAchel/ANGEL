package sqlinject

import (
	"time"
)

type Sqlinject0164 struct{}

func NewSqlinject0164() *Sqlinject0164 {
	return &Sqlinject0164{}
}

func (e *Sqlinject0164) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0164) Name() string { return "Sqlinject0164" }
func (e *Sqlinject0164) Timestamp() time.Time { return time.Now() }
