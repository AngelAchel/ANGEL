package sqlinject

import (
	"time"
)

type Sqlinject0090 struct{}

func NewSqlinject0090() *Sqlinject0090 {
	return &Sqlinject0090{}
}

func (e *Sqlinject0090) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0090) Name() string { return "Sqlinject0090" }
func (e *Sqlinject0090) Timestamp() time.Time { return time.Now() }
