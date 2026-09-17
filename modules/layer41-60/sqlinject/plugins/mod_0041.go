package sqlinject

import (
	"time"
)

type Sqlinject0041 struct{}

func NewSqlinject0041() *Sqlinject0041 {
	return &Sqlinject0041{}
}

func (e *Sqlinject0041) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0041) Name() string { return "Sqlinject0041" }
func (e *Sqlinject0041) Timestamp() time.Time { return time.Now() }
