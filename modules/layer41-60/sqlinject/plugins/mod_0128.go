package sqlinject

import (
	"time"
)

type Sqlinject0128 struct{}

func NewSqlinject0128() *Sqlinject0128 {
	return &Sqlinject0128{}
}

func (e *Sqlinject0128) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0128) Name() string { return "Sqlinject0128" }
func (e *Sqlinject0128) Timestamp() time.Time { return time.Now() }
