package sqlinject

import (
	"time"
)

type Sqlinject0075 struct{}

func NewSqlinject0075() *Sqlinject0075 {
	return &Sqlinject0075{}
}

func (e *Sqlinject0075) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0075) Name() string { return "Sqlinject0075" }
func (e *Sqlinject0075) Timestamp() time.Time { return time.Now() }
