package sqlinject

import (
	"time"
)

type Sqlinject0013 struct{}

func NewSqlinject0013() *Sqlinject0013 {
	return &Sqlinject0013{}
}

func (e *Sqlinject0013) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0013) Name() string { return "Sqlinject0013" }
func (e *Sqlinject0013) Timestamp() time.Time { return time.Now() }
