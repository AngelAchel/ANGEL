package sqlinject

import (
	"time"
)

type Sqlinject0023 struct{}

func NewSqlinject0023() *Sqlinject0023 {
	return &Sqlinject0023{}
}

func (e *Sqlinject0023) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0023) Name() string { return "Sqlinject0023" }
func (e *Sqlinject0023) Timestamp() time.Time { return time.Now() }
