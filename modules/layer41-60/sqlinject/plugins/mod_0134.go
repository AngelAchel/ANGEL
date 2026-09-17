package sqlinject

import (
	"time"
)

type Sqlinject0134 struct{}

func NewSqlinject0134() *Sqlinject0134 {
	return &Sqlinject0134{}
}

func (e *Sqlinject0134) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0134) Name() string { return "Sqlinject0134" }
func (e *Sqlinject0134) Timestamp() time.Time { return time.Now() }
