package sqlinject

import (
	"time"
)

type Sqlinject0120 struct{}

func NewSqlinject0120() *Sqlinject0120 {
	return &Sqlinject0120{}
}

func (e *Sqlinject0120) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0120) Name() string { return "Sqlinject0120" }
func (e *Sqlinject0120) Timestamp() time.Time { return time.Now() }
