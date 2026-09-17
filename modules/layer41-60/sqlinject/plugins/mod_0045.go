package sqlinject

import (
	"time"
)

type Sqlinject0045 struct{}

func NewSqlinject0045() *Sqlinject0045 {
	return &Sqlinject0045{}
}

func (e *Sqlinject0045) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0045) Name() string { return "Sqlinject0045" }
func (e *Sqlinject0045) Timestamp() time.Time { return time.Now() }
