package sqlinject

import (
	"time"
)

type Sqlinject0020 struct{}

func NewSqlinject0020() *Sqlinject0020 {
	return &Sqlinject0020{}
}

func (e *Sqlinject0020) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0020) Name() string { return "Sqlinject0020" }
func (e *Sqlinject0020) Timestamp() time.Time { return time.Now() }
