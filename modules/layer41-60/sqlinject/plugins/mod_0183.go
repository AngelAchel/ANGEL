package sqlinject

import (
	"time"
)

type Sqlinject0183 struct{}

func NewSqlinject0183() *Sqlinject0183 {
	return &Sqlinject0183{}
}

func (e *Sqlinject0183) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0183) Name() string { return "Sqlinject0183" }
func (e *Sqlinject0183) Timestamp() time.Time { return time.Now() }
