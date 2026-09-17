package sqlinject

import (
	"time"
)

type Sqlinject0147 struct{}

func NewSqlinject0147() *Sqlinject0147 {
	return &Sqlinject0147{}
}

func (e *Sqlinject0147) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0147) Name() string { return "Sqlinject0147" }
func (e *Sqlinject0147) Timestamp() time.Time { return time.Now() }
