package sqlinject

import (
	"time"
)

type Sqlinject0116 struct{}

func NewSqlinject0116() *Sqlinject0116 {
	return &Sqlinject0116{}
}

func (e *Sqlinject0116) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0116) Name() string { return "Sqlinject0116" }
func (e *Sqlinject0116) Timestamp() time.Time { return time.Now() }
