package sqlinject

import (
	"time"
)

type Sqlinject0015 struct{}

func NewSqlinject0015() *Sqlinject0015 {
	return &Sqlinject0015{}
}

func (e *Sqlinject0015) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0015) Name() string { return "Sqlinject0015" }
func (e *Sqlinject0015) Timestamp() time.Time { return time.Now() }
