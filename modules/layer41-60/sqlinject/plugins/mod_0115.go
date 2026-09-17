package sqlinject

import (
	"time"
)

type Sqlinject0115 struct{}

func NewSqlinject0115() *Sqlinject0115 {
	return &Sqlinject0115{}
}

func (e *Sqlinject0115) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0115) Name() string { return "Sqlinject0115" }
func (e *Sqlinject0115) Timestamp() time.Time { return time.Now() }
