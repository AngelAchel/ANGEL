package sqlinject

import (
	"time"
)

type Sqlinject0087 struct{}

func NewSqlinject0087() *Sqlinject0087 {
	return &Sqlinject0087{}
}

func (e *Sqlinject0087) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0087) Name() string { return "Sqlinject0087" }
func (e *Sqlinject0087) Timestamp() time.Time { return time.Now() }
