package sqlinject

import (
	"time"
)

type Sqlinject0088 struct{}

func NewSqlinject0088() *Sqlinject0088 {
	return &Sqlinject0088{}
}

func (e *Sqlinject0088) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0088) Name() string { return "Sqlinject0088" }
func (e *Sqlinject0088) Timestamp() time.Time { return time.Now() }
