package sqlinject

import (
	"time"
)

type Sqlinject0082 struct{}

func NewSqlinject0082() *Sqlinject0082 {
	return &Sqlinject0082{}
}

func (e *Sqlinject0082) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0082) Name() string { return "Sqlinject0082" }
func (e *Sqlinject0082) Timestamp() time.Time { return time.Now() }
