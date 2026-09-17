package sqlinject

import (
	"time"
)

type Sqlinject0081 struct{}

func NewSqlinject0081() *Sqlinject0081 {
	return &Sqlinject0081{}
}

func (e *Sqlinject0081) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0081) Name() string { return "Sqlinject0081" }
func (e *Sqlinject0081) Timestamp() time.Time { return time.Now() }
