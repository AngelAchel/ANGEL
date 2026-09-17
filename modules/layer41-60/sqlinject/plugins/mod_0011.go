package sqlinject

import (
	"time"
)

type Sqlinject0011 struct{}

func NewSqlinject0011() *Sqlinject0011 {
	return &Sqlinject0011{}
}

func (e *Sqlinject0011) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0011) Name() string { return "Sqlinject0011" }
func (e *Sqlinject0011) Timestamp() time.Time { return time.Now() }
