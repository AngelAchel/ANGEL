package sqlinject

import (
	"time"
)

type Sqlinject0085 struct{}

func NewSqlinject0085() *Sqlinject0085 {
	return &Sqlinject0085{}
}

func (e *Sqlinject0085) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0085) Name() string { return "Sqlinject0085" }
func (e *Sqlinject0085) Timestamp() time.Time { return time.Now() }
