package sqlinject

import (
	"time"
)

type Sqlinject0008 struct{}

func NewSqlinject0008() *Sqlinject0008 {
	return &Sqlinject0008{}
}

func (e *Sqlinject0008) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0008) Name() string { return "Sqlinject0008" }
func (e *Sqlinject0008) Timestamp() time.Time { return time.Now() }
