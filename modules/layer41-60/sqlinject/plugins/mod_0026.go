package sqlinject

import (
	"time"
)

type Sqlinject0026 struct{}

func NewSqlinject0026() *Sqlinject0026 {
	return &Sqlinject0026{}
}

func (e *Sqlinject0026) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0026) Name() string { return "Sqlinject0026" }
func (e *Sqlinject0026) Timestamp() time.Time { return time.Now() }
