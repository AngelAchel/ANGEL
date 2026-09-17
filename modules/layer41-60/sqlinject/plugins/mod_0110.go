package sqlinject

import (
	"time"
)

type Sqlinject0110 struct{}

func NewSqlinject0110() *Sqlinject0110 {
	return &Sqlinject0110{}
}

func (e *Sqlinject0110) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0110) Name() string { return "Sqlinject0110" }
func (e *Sqlinject0110) Timestamp() time.Time { return time.Now() }
