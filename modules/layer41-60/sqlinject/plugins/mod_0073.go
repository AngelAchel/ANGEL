package sqlinject

import (
	"time"
)

type Sqlinject0073 struct{}

func NewSqlinject0073() *Sqlinject0073 {
	return &Sqlinject0073{}
}

func (e *Sqlinject0073) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0073) Name() string { return "Sqlinject0073" }
func (e *Sqlinject0073) Timestamp() time.Time { return time.Now() }
