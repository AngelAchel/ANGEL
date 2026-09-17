package sqlinject

import (
	"time"
)

type Sqlinject0006 struct{}

func NewSqlinject0006() *Sqlinject0006 {
	return &Sqlinject0006{}
}

func (e *Sqlinject0006) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0006) Name() string { return "Sqlinject0006" }
func (e *Sqlinject0006) Timestamp() time.Time { return time.Now() }
