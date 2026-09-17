package sqlinject

import (
	"time"
)

type Sqlinject0018 struct{}

func NewSqlinject0018() *Sqlinject0018 {
	return &Sqlinject0018{}
}

func (e *Sqlinject0018) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0018) Name() string { return "Sqlinject0018" }
func (e *Sqlinject0018) Timestamp() time.Time { return time.Now() }
