package sqlinject

import (
	"time"
)

type Sqlinject0171 struct{}

func NewSqlinject0171() *Sqlinject0171 {
	return &Sqlinject0171{}
}

func (e *Sqlinject0171) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0171) Name() string { return "Sqlinject0171" }
func (e *Sqlinject0171) Timestamp() time.Time { return time.Now() }
