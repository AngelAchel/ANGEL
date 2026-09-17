package sqlinject

import (
	"time"
)

type Sqlinject0037 struct{}

func NewSqlinject0037() *Sqlinject0037 {
	return &Sqlinject0037{}
}

func (e *Sqlinject0037) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0037) Name() string { return "Sqlinject0037" }
func (e *Sqlinject0037) Timestamp() time.Time { return time.Now() }
