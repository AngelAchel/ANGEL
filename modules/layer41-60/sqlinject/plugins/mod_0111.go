package sqlinject

import (
	"time"
)

type Sqlinject0111 struct{}

func NewSqlinject0111() *Sqlinject0111 {
	return &Sqlinject0111{}
}

func (e *Sqlinject0111) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0111) Name() string { return "Sqlinject0111" }
func (e *Sqlinject0111) Timestamp() time.Time { return time.Now() }
