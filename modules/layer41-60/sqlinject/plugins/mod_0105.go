package sqlinject

import (
	"time"
)

type Sqlinject0105 struct{}

func NewSqlinject0105() *Sqlinject0105 {
	return &Sqlinject0105{}
}

func (e *Sqlinject0105) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0105) Name() string { return "Sqlinject0105" }
func (e *Sqlinject0105) Timestamp() time.Time { return time.Now() }
