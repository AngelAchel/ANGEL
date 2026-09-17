package sqlinject

import (
	"time"
)

type Sqlinject0141 struct{}

func NewSqlinject0141() *Sqlinject0141 {
	return &Sqlinject0141{}
}

func (e *Sqlinject0141) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0141) Name() string { return "Sqlinject0141" }
func (e *Sqlinject0141) Timestamp() time.Time { return time.Now() }
