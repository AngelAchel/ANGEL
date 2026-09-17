package sqlinject

import (
	"time"
)

type Sqlinject0101 struct{}

func NewSqlinject0101() *Sqlinject0101 {
	return &Sqlinject0101{}
}

func (e *Sqlinject0101) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0101) Name() string { return "Sqlinject0101" }
func (e *Sqlinject0101) Timestamp() time.Time { return time.Now() }
