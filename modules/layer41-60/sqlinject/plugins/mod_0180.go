package sqlinject

import (
	"time"
)

type Sqlinject0180 struct{}

func NewSqlinject0180() *Sqlinject0180 {
	return &Sqlinject0180{}
}

func (e *Sqlinject0180) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0180) Name() string { return "Sqlinject0180" }
func (e *Sqlinject0180) Timestamp() time.Time { return time.Now() }
