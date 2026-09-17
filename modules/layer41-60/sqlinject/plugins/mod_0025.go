package sqlinject

import (
	"time"
)

type Sqlinject0025 struct{}

func NewSqlinject0025() *Sqlinject0025 {
	return &Sqlinject0025{}
}

func (e *Sqlinject0025) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0025) Name() string { return "Sqlinject0025" }
func (e *Sqlinject0025) Timestamp() time.Time { return time.Now() }
