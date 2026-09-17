package sqlinject

import (
	"time"
)

type Sqlinject0125 struct{}

func NewSqlinject0125() *Sqlinject0125 {
	return &Sqlinject0125{}
}

func (e *Sqlinject0125) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0125) Name() string { return "Sqlinject0125" }
func (e *Sqlinject0125) Timestamp() time.Time { return time.Now() }
