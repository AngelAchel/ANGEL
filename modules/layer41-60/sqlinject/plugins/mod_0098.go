package sqlinject

import (
	"time"
)

type Sqlinject0098 struct{}

func NewSqlinject0098() *Sqlinject0098 {
	return &Sqlinject0098{}
}

func (e *Sqlinject0098) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0098) Name() string { return "Sqlinject0098" }
func (e *Sqlinject0098) Timestamp() time.Time { return time.Now() }
