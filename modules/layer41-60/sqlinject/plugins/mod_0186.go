package sqlinject

import (
	"time"
)

type Sqlinject0186 struct{}

func NewSqlinject0186() *Sqlinject0186 {
	return &Sqlinject0186{}
}

func (e *Sqlinject0186) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0186) Name() string { return "Sqlinject0186" }
func (e *Sqlinject0186) Timestamp() time.Time { return time.Now() }
