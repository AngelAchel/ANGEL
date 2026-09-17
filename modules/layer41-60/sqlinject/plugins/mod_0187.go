package sqlinject

import (
	"time"
)

type Sqlinject0187 struct{}

func NewSqlinject0187() *Sqlinject0187 {
	return &Sqlinject0187{}
}

func (e *Sqlinject0187) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0187) Name() string { return "Sqlinject0187" }
func (e *Sqlinject0187) Timestamp() time.Time { return time.Now() }
