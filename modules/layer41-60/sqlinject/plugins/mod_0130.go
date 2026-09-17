package sqlinject

import (
	"time"
)

type Sqlinject0130 struct{}

func NewSqlinject0130() *Sqlinject0130 {
	return &Sqlinject0130{}
}

func (e *Sqlinject0130) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0130) Name() string { return "Sqlinject0130" }
func (e *Sqlinject0130) Timestamp() time.Time { return time.Now() }
