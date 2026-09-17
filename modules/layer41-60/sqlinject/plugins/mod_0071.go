package sqlinject

import (
	"time"
)

type Sqlinject0071 struct{}

func NewSqlinject0071() *Sqlinject0071 {
	return &Sqlinject0071{}
}

func (e *Sqlinject0071) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0071) Name() string { return "Sqlinject0071" }
func (e *Sqlinject0071) Timestamp() time.Time { return time.Now() }
