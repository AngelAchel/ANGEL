package sqlinject

import (
	"time"
)

type Sqlinject0135 struct{}

func NewSqlinject0135() *Sqlinject0135 {
	return &Sqlinject0135{}
}

func (e *Sqlinject0135) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0135) Name() string { return "Sqlinject0135" }
func (e *Sqlinject0135) Timestamp() time.Time { return time.Now() }
