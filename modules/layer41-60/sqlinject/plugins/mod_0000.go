package sqlinject

import (
	"time"
)

type Sqlinject0000 struct{}

func NewSqlinject0000() *Sqlinject0000 {
	return &Sqlinject0000{}
}

func (e *Sqlinject0000) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0000) Name() string { return "Sqlinject0000" }
func (e *Sqlinject0000) Timestamp() time.Time { return time.Now() }
