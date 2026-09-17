package sqlinject

import (
	"time"
)

type Sqlinject0194 struct{}

func NewSqlinject0194() *Sqlinject0194 {
	return &Sqlinject0194{}
}

func (e *Sqlinject0194) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0194) Name() string { return "Sqlinject0194" }
func (e *Sqlinject0194) Timestamp() time.Time { return time.Now() }
