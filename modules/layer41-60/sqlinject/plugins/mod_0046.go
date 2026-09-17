package sqlinject

import (
	"time"
)

type Sqlinject0046 struct{}

func NewSqlinject0046() *Sqlinject0046 {
	return &Sqlinject0046{}
}

func (e *Sqlinject0046) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0046) Name() string { return "Sqlinject0046" }
func (e *Sqlinject0046) Timestamp() time.Time { return time.Now() }
