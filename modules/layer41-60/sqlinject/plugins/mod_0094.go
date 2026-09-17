package sqlinject

import (
	"time"
)

type Sqlinject0094 struct{}

func NewSqlinject0094() *Sqlinject0094 {
	return &Sqlinject0094{}
}

func (e *Sqlinject0094) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0094) Name() string { return "Sqlinject0094" }
func (e *Sqlinject0094) Timestamp() time.Time { return time.Now() }
