package sqlinject

import (
	"time"
)

type Sqlinject0003 struct{}

func NewSqlinject0003() *Sqlinject0003 {
	return &Sqlinject0003{}
}

func (e *Sqlinject0003) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0003) Name() string { return "Sqlinject0003" }
func (e *Sqlinject0003) Timestamp() time.Time { return time.Now() }
