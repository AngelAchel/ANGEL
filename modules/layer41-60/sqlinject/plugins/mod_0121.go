package sqlinject

import (
	"time"
)

type Sqlinject0121 struct{}

func NewSqlinject0121() *Sqlinject0121 {
	return &Sqlinject0121{}
}

func (e *Sqlinject0121) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0121) Name() string { return "Sqlinject0121" }
func (e *Sqlinject0121) Timestamp() time.Time { return time.Now() }
