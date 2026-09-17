package sqlinject

import (
	"time"
)

type Sqlinject0182 struct{}

func NewSqlinject0182() *Sqlinject0182 {
	return &Sqlinject0182{}
}

func (e *Sqlinject0182) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0182) Name() string { return "Sqlinject0182" }
func (e *Sqlinject0182) Timestamp() time.Time { return time.Now() }
