package sqlinject

import (
	"time"
)

type Sqlinject0189 struct{}

func NewSqlinject0189() *Sqlinject0189 {
	return &Sqlinject0189{}
}

func (e *Sqlinject0189) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0189) Name() string { return "Sqlinject0189" }
func (e *Sqlinject0189) Timestamp() time.Time { return time.Now() }
