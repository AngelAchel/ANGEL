package sqlinject

import (
	"time"
)

type Sqlinject0172 struct{}

func NewSqlinject0172() *Sqlinject0172 {
	return &Sqlinject0172{}
}

func (e *Sqlinject0172) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0172) Name() string { return "Sqlinject0172" }
func (e *Sqlinject0172) Timestamp() time.Time { return time.Now() }
