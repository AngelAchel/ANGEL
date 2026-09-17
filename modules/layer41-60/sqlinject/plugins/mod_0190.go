package sqlinject

import (
	"time"
)

type Sqlinject0190 struct{}

func NewSqlinject0190() *Sqlinject0190 {
	return &Sqlinject0190{}
}

func (e *Sqlinject0190) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0190) Name() string { return "Sqlinject0190" }
func (e *Sqlinject0190) Timestamp() time.Time { return time.Now() }
