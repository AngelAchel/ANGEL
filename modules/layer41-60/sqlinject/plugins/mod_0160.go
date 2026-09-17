package sqlinject

import (
	"time"
)

type Sqlinject0160 struct{}

func NewSqlinject0160() *Sqlinject0160 {
	return &Sqlinject0160{}
}

func (e *Sqlinject0160) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0160) Name() string { return "Sqlinject0160" }
func (e *Sqlinject0160) Timestamp() time.Time { return time.Now() }
