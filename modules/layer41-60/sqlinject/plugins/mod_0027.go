package sqlinject

import (
	"time"
)

type Sqlinject0027 struct{}

func NewSqlinject0027() *Sqlinject0027 {
	return &Sqlinject0027{}
}

func (e *Sqlinject0027) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0027) Name() string { return "Sqlinject0027" }
func (e *Sqlinject0027) Timestamp() time.Time { return time.Now() }
