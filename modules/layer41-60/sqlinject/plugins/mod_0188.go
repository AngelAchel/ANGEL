package sqlinject

import (
	"time"
)

type Sqlinject0188 struct{}

func NewSqlinject0188() *Sqlinject0188 {
	return &Sqlinject0188{}
}

func (e *Sqlinject0188) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0188) Name() string { return "Sqlinject0188" }
func (e *Sqlinject0188) Timestamp() time.Time { return time.Now() }
