package sqlinject

import (
	"time"
)

type Sqlinject0038 struct{}

func NewSqlinject0038() *Sqlinject0038 {
	return &Sqlinject0038{}
}

func (e *Sqlinject0038) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0038) Name() string { return "Sqlinject0038" }
func (e *Sqlinject0038) Timestamp() time.Time { return time.Now() }
