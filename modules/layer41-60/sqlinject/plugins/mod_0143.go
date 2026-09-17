package sqlinject

import (
	"time"
)

type Sqlinject0143 struct{}

func NewSqlinject0143() *Sqlinject0143 {
	return &Sqlinject0143{}
}

func (e *Sqlinject0143) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0143) Name() string { return "Sqlinject0143" }
func (e *Sqlinject0143) Timestamp() time.Time { return time.Now() }
