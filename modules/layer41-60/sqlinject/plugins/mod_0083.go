package sqlinject

import (
	"time"
)

type Sqlinject0083 struct{}

func NewSqlinject0083() *Sqlinject0083 {
	return &Sqlinject0083{}
}

func (e *Sqlinject0083) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0083) Name() string { return "Sqlinject0083" }
func (e *Sqlinject0083) Timestamp() time.Time { return time.Now() }
