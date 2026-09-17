package sqlinject

import (
	"time"
)

type Sqlinject0010 struct{}

func NewSqlinject0010() *Sqlinject0010 {
	return &Sqlinject0010{}
}

func (e *Sqlinject0010) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0010) Name() string { return "Sqlinject0010" }
func (e *Sqlinject0010) Timestamp() time.Time { return time.Now() }
