package sqlinject

import (
	"time"
)

type Sqlinject0181 struct{}

func NewSqlinject0181() *Sqlinject0181 {
	return &Sqlinject0181{}
}

func (e *Sqlinject0181) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0181) Name() string { return "Sqlinject0181" }
func (e *Sqlinject0181) Timestamp() time.Time { return time.Now() }
