package sqlinject

import (
	"time"
)

type Sqlinject0117 struct{}

func NewSqlinject0117() *Sqlinject0117 {
	return &Sqlinject0117{}
}

func (e *Sqlinject0117) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0117) Name() string { return "Sqlinject0117" }
func (e *Sqlinject0117) Timestamp() time.Time { return time.Now() }
