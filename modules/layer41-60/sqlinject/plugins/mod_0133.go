package sqlinject

import (
	"time"
)

type Sqlinject0133 struct{}

func NewSqlinject0133() *Sqlinject0133 {
	return &Sqlinject0133{}
}

func (e *Sqlinject0133) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0133) Name() string { return "Sqlinject0133" }
func (e *Sqlinject0133) Timestamp() time.Time { return time.Now() }
