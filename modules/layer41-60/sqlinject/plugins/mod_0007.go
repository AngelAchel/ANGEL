package sqlinject

import (
	"time"
)

type Sqlinject0007 struct{}

func NewSqlinject0007() *Sqlinject0007 {
	return &Sqlinject0007{}
}

func (e *Sqlinject0007) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0007) Name() string { return "Sqlinject0007" }
func (e *Sqlinject0007) Timestamp() time.Time { return time.Now() }
