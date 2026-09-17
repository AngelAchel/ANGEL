package sqlinject

import (
	"time"
)

type Sqlinject0122 struct{}

func NewSqlinject0122() *Sqlinject0122 {
	return &Sqlinject0122{}
}

func (e *Sqlinject0122) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0122) Name() string { return "Sqlinject0122" }
func (e *Sqlinject0122) Timestamp() time.Time { return time.Now() }
