package sqlinject

import (
	"time"
)

type Sqlinject0103 struct{}

func NewSqlinject0103() *Sqlinject0103 {
	return &Sqlinject0103{}
}

func (e *Sqlinject0103) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0103) Name() string { return "Sqlinject0103" }
func (e *Sqlinject0103) Timestamp() time.Time { return time.Now() }
