package sqlinject

import (
	"time"
)

type Sqlinject0100 struct{}

func NewSqlinject0100() *Sqlinject0100 {
	return &Sqlinject0100{}
}

func (e *Sqlinject0100) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0100) Name() string { return "Sqlinject0100" }
func (e *Sqlinject0100) Timestamp() time.Time { return time.Now() }
