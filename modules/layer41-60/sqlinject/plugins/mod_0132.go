package sqlinject

import (
	"time"
)

type Sqlinject0132 struct{}

func NewSqlinject0132() *Sqlinject0132 {
	return &Sqlinject0132{}
}

func (e *Sqlinject0132) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0132) Name() string { return "Sqlinject0132" }
func (e *Sqlinject0132) Timestamp() time.Time { return time.Now() }
