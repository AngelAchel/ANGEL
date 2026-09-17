package sqlinject

import (
	"time"
)

type Sqlinject0005 struct{}

func NewSqlinject0005() *Sqlinject0005 {
	return &Sqlinject0005{}
}

func (e *Sqlinject0005) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0005) Name() string { return "Sqlinject0005" }
func (e *Sqlinject0005) Timestamp() time.Time { return time.Now() }
