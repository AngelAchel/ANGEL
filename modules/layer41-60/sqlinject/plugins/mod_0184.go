package sqlinject

import (
	"time"
)

type Sqlinject0184 struct{}

func NewSqlinject0184() *Sqlinject0184 {
	return &Sqlinject0184{}
}

func (e *Sqlinject0184) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0184) Name() string { return "Sqlinject0184" }
func (e *Sqlinject0184) Timestamp() time.Time { return time.Now() }
