package sqlinject

import (
	"time"
)

type Sqlinject0175 struct{}

func NewSqlinject0175() *Sqlinject0175 {
	return &Sqlinject0175{}
}

func (e *Sqlinject0175) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0175) Name() string { return "Sqlinject0175" }
func (e *Sqlinject0175) Timestamp() time.Time { return time.Now() }
