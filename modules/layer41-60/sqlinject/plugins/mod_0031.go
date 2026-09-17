package sqlinject

import (
	"time"
)

type Sqlinject0031 struct{}

func NewSqlinject0031() *Sqlinject0031 {
	return &Sqlinject0031{}
}

func (e *Sqlinject0031) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0031) Name() string { return "Sqlinject0031" }
func (e *Sqlinject0031) Timestamp() time.Time { return time.Now() }
