package sqlinject

import (
	"time"
)

type Sqlinject0198 struct{}

func NewSqlinject0198() *Sqlinject0198 {
	return &Sqlinject0198{}
}

func (e *Sqlinject0198) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0198) Name() string { return "Sqlinject0198" }
func (e *Sqlinject0198) Timestamp() time.Time { return time.Now() }
