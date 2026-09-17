package sqlinject

import (
	"time"
)

type Sqlinject0086 struct{}

func NewSqlinject0086() *Sqlinject0086 {
	return &Sqlinject0086{}
}

func (e *Sqlinject0086) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0086) Name() string { return "Sqlinject0086" }
func (e *Sqlinject0086) Timestamp() time.Time { return time.Now() }
