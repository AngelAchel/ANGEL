package sqlinject

import (
	"time"
)

type Sqlinject0001 struct{}

func NewSqlinject0001() *Sqlinject0001 {
	return &Sqlinject0001{}
}

func (e *Sqlinject0001) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0001) Name() string { return "Sqlinject0001" }
func (e *Sqlinject0001) Timestamp() time.Time { return time.Now() }
