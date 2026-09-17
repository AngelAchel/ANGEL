package sqlinject

import (
	"time"
)

type Sqlinject0002 struct{}

func NewSqlinject0002() *Sqlinject0002 {
	return &Sqlinject0002{}
}

func (e *Sqlinject0002) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0002) Name() string { return "Sqlinject0002" }
func (e *Sqlinject0002) Timestamp() time.Time { return time.Now() }
