package sqlinject

import (
	"time"
)

type Sqlinject0119 struct{}

func NewSqlinject0119() *Sqlinject0119 {
	return &Sqlinject0119{}
}

func (e *Sqlinject0119) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0119) Name() string { return "Sqlinject0119" }
func (e *Sqlinject0119) Timestamp() time.Time { return time.Now() }
