package sqlinject

import (
	"time"
)

type Sqlinject0151 struct{}

func NewSqlinject0151() *Sqlinject0151 {
	return &Sqlinject0151{}
}

func (e *Sqlinject0151) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0151) Name() string { return "Sqlinject0151" }
func (e *Sqlinject0151) Timestamp() time.Time { return time.Now() }
