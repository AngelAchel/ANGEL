package sqlinject

import (
	"time"
)

type Sqlinject0139 struct{}

func NewSqlinject0139() *Sqlinject0139 {
	return &Sqlinject0139{}
}

func (e *Sqlinject0139) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0139) Name() string { return "Sqlinject0139" }
func (e *Sqlinject0139) Timestamp() time.Time { return time.Now() }
