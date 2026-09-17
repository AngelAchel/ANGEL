package sqlinject

import (
	"time"
)

type Sqlinject0077 struct{}

func NewSqlinject0077() *Sqlinject0077 {
	return &Sqlinject0077{}
}

func (e *Sqlinject0077) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0077) Name() string { return "Sqlinject0077" }
func (e *Sqlinject0077) Timestamp() time.Time { return time.Now() }
