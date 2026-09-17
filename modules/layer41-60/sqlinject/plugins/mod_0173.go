package sqlinject

import (
	"time"
)

type Sqlinject0173 struct{}

func NewSqlinject0173() *Sqlinject0173 {
	return &Sqlinject0173{}
}

func (e *Sqlinject0173) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0173) Name() string { return "Sqlinject0173" }
func (e *Sqlinject0173) Timestamp() time.Time { return time.Now() }
