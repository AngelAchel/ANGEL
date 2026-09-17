package sqlinject

import (
	"time"
)

type Sqlinject0089 struct{}

func NewSqlinject0089() *Sqlinject0089 {
	return &Sqlinject0089{}
}

func (e *Sqlinject0089) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0089) Name() string { return "Sqlinject0089" }
func (e *Sqlinject0089) Timestamp() time.Time { return time.Now() }
