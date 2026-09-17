package sqlinject

import (
	"time"
)

type Sqlinject0084 struct{}

func NewSqlinject0084() *Sqlinject0084 {
	return &Sqlinject0084{}
}

func (e *Sqlinject0084) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0084) Name() string { return "Sqlinject0084" }
func (e *Sqlinject0084) Timestamp() time.Time { return time.Now() }
