package sqlinject

import (
	"time"
)

type Sqlinject0113 struct{}

func NewSqlinject0113() *Sqlinject0113 {
	return &Sqlinject0113{}
}

func (e *Sqlinject0113) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0113) Name() string { return "Sqlinject0113" }
func (e *Sqlinject0113) Timestamp() time.Time { return time.Now() }
