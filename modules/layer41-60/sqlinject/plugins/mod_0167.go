package sqlinject

import (
	"time"
)

type Sqlinject0167 struct{}

func NewSqlinject0167() *Sqlinject0167 {
	return &Sqlinject0167{}
}

func (e *Sqlinject0167) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0167) Name() string { return "Sqlinject0167" }
func (e *Sqlinject0167) Timestamp() time.Time { return time.Now() }
