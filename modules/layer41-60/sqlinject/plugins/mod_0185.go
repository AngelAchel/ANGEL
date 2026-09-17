package sqlinject

import (
	"time"
)

type Sqlinject0185 struct{}

func NewSqlinject0185() *Sqlinject0185 {
	return &Sqlinject0185{}
}

func (e *Sqlinject0185) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0185) Name() string { return "Sqlinject0185" }
func (e *Sqlinject0185) Timestamp() time.Time { return time.Now() }
