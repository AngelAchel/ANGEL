package sqlinject

import (
	"time"
)

type Sqlinject0140 struct{}

func NewSqlinject0140() *Sqlinject0140 {
	return &Sqlinject0140{}
}

func (e *Sqlinject0140) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0140) Name() string { return "Sqlinject0140" }
func (e *Sqlinject0140) Timestamp() time.Time { return time.Now() }
