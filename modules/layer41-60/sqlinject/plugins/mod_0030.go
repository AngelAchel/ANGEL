package sqlinject

import (
	"time"
)

type Sqlinject0030 struct{}

func NewSqlinject0030() *Sqlinject0030 {
	return &Sqlinject0030{}
}

func (e *Sqlinject0030) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0030) Name() string { return "Sqlinject0030" }
func (e *Sqlinject0030) Timestamp() time.Time { return time.Now() }
