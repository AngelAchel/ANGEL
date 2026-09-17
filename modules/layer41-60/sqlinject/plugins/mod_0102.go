package sqlinject

import (
	"time"
)

type Sqlinject0102 struct{}

func NewSqlinject0102() *Sqlinject0102 {
	return &Sqlinject0102{}
}

func (e *Sqlinject0102) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0102) Name() string { return "Sqlinject0102" }
func (e *Sqlinject0102) Timestamp() time.Time { return time.Now() }
