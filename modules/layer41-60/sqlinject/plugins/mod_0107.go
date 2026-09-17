package sqlinject

import (
	"time"
)

type Sqlinject0107 struct{}

func NewSqlinject0107() *Sqlinject0107 {
	return &Sqlinject0107{}
}

func (e *Sqlinject0107) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0107) Name() string { return "Sqlinject0107" }
func (e *Sqlinject0107) Timestamp() time.Time { return time.Now() }
