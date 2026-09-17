package sqlinject

import (
	"time"
)

type Sqlinject0177 struct{}

func NewSqlinject0177() *Sqlinject0177 {
	return &Sqlinject0177{}
}

func (e *Sqlinject0177) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0177) Name() string { return "Sqlinject0177" }
func (e *Sqlinject0177) Timestamp() time.Time { return time.Now() }
