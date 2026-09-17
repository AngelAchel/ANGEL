package sqlinject

import (
	"time"
)

type Sqlinject0004 struct{}

func NewSqlinject0004() *Sqlinject0004 {
	return &Sqlinject0004{}
}

func (e *Sqlinject0004) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0004) Name() string { return "Sqlinject0004" }
func (e *Sqlinject0004) Timestamp() time.Time { return time.Now() }
