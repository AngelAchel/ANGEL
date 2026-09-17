package sqlinject

import (
	"time"
)

type Sqlinject0060 struct{}

func NewSqlinject0060() *Sqlinject0060 {
	return &Sqlinject0060{}
}

func (e *Sqlinject0060) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0060) Name() string { return "Sqlinject0060" }
func (e *Sqlinject0060) Timestamp() time.Time { return time.Now() }
