package sqlinject

import (
	"time"
)

type Sqlinject0028 struct{}

func NewSqlinject0028() *Sqlinject0028 {
	return &Sqlinject0028{}
}

func (e *Sqlinject0028) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0028) Name() string { return "Sqlinject0028" }
func (e *Sqlinject0028) Timestamp() time.Time { return time.Now() }
