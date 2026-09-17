package sqlinject

import (
	"time"
)

type Sqlinject0127 struct{}

func NewSqlinject0127() *Sqlinject0127 {
	return &Sqlinject0127{}
}

func (e *Sqlinject0127) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0127) Name() string { return "Sqlinject0127" }
func (e *Sqlinject0127) Timestamp() time.Time { return time.Now() }
