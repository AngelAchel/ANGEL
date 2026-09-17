package sqlinject

import (
	"time"
)

type Sqlinject0174 struct{}

func NewSqlinject0174() *Sqlinject0174 {
	return &Sqlinject0174{}
}

func (e *Sqlinject0174) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0174) Name() string { return "Sqlinject0174" }
func (e *Sqlinject0174) Timestamp() time.Time { return time.Now() }
