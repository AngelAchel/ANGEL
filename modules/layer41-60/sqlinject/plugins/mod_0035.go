package sqlinject

import (
	"time"
)

type Sqlinject0035 struct{}

func NewSqlinject0035() *Sqlinject0035 {
	return &Sqlinject0035{}
}

func (e *Sqlinject0035) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0035) Name() string { return "Sqlinject0035" }
func (e *Sqlinject0035) Timestamp() time.Time { return time.Now() }
