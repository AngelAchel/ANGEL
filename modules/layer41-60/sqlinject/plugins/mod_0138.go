package sqlinject

import (
	"time"
)

type Sqlinject0138 struct{}

func NewSqlinject0138() *Sqlinject0138 {
	return &Sqlinject0138{}
}

func (e *Sqlinject0138) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0138) Name() string { return "Sqlinject0138" }
func (e *Sqlinject0138) Timestamp() time.Time { return time.Now() }
