package sqlinject

import (
	"time"
)

type Sqlinject0161 struct{}

func NewSqlinject0161() *Sqlinject0161 {
	return &Sqlinject0161{}
}

func (e *Sqlinject0161) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0161) Name() string { return "Sqlinject0161" }
func (e *Sqlinject0161) Timestamp() time.Time { return time.Now() }
