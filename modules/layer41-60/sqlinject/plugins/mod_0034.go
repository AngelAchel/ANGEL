package sqlinject

import (
	"time"
)

type Sqlinject0034 struct{}

func NewSqlinject0034() *Sqlinject0034 {
	return &Sqlinject0034{}
}

func (e *Sqlinject0034) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0034) Name() string { return "Sqlinject0034" }
func (e *Sqlinject0034) Timestamp() time.Time { return time.Now() }
