package sqlinject

import (
	"time"
)

type Sqlinject0062 struct{}

func NewSqlinject0062() *Sqlinject0062 {
	return &Sqlinject0062{}
}

func (e *Sqlinject0062) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0062) Name() string { return "Sqlinject0062" }
func (e *Sqlinject0062) Timestamp() time.Time { return time.Now() }
