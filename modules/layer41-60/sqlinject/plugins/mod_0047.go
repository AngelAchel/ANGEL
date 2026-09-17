package sqlinject

import (
	"time"
)

type Sqlinject0047 struct{}

func NewSqlinject0047() *Sqlinject0047 {
	return &Sqlinject0047{}
}

func (e *Sqlinject0047) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0047) Name() string { return "Sqlinject0047" }
func (e *Sqlinject0047) Timestamp() time.Time { return time.Now() }
