package sqlinject

import (
	"time"
)

type Sqlinject0145 struct{}

func NewSqlinject0145() *Sqlinject0145 {
	return &Sqlinject0145{}
}

func (e *Sqlinject0145) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0145) Name() string { return "Sqlinject0145" }
func (e *Sqlinject0145) Timestamp() time.Time { return time.Now() }
