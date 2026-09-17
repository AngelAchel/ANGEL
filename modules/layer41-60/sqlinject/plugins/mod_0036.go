package sqlinject

import (
	"time"
)

type Sqlinject0036 struct{}

func NewSqlinject0036() *Sqlinject0036 {
	return &Sqlinject0036{}
}

func (e *Sqlinject0036) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0036) Name() string { return "Sqlinject0036" }
func (e *Sqlinject0036) Timestamp() time.Time { return time.Now() }
