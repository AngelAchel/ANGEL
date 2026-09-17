package sqlinject

import (
	"time"
)

type Sqlinject0170 struct{}

func NewSqlinject0170() *Sqlinject0170 {
	return &Sqlinject0170{}
}

func (e *Sqlinject0170) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0170) Name() string { return "Sqlinject0170" }
func (e *Sqlinject0170) Timestamp() time.Time { return time.Now() }
