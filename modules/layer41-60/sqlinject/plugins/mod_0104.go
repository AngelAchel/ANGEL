package sqlinject

import (
	"time"
)

type Sqlinject0104 struct{}

func NewSqlinject0104() *Sqlinject0104 {
	return &Sqlinject0104{}
}

func (e *Sqlinject0104) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0104) Name() string { return "Sqlinject0104" }
func (e *Sqlinject0104) Timestamp() time.Time { return time.Now() }
