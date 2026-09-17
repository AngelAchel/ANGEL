package sqlinject

import (
	"time"
)

type Sqlinject0012 struct{}

func NewSqlinject0012() *Sqlinject0012 {
	return &Sqlinject0012{}
}

func (e *Sqlinject0012) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0012) Name() string { return "Sqlinject0012" }
func (e *Sqlinject0012) Timestamp() time.Time { return time.Now() }
