package sqlinject

import (
	"time"
)

type Sqlinject0096 struct{}

func NewSqlinject0096() *Sqlinject0096 {
	return &Sqlinject0096{}
}

func (e *Sqlinject0096) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0096) Name() string { return "Sqlinject0096" }
func (e *Sqlinject0096) Timestamp() time.Time { return time.Now() }
