package sqlinject

import (
	"time"
)

type Sqlinject0061 struct{}

func NewSqlinject0061() *Sqlinject0061 {
	return &Sqlinject0061{}
}

func (e *Sqlinject0061) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0061) Name() string { return "Sqlinject0061" }
func (e *Sqlinject0061) Timestamp() time.Time { return time.Now() }
