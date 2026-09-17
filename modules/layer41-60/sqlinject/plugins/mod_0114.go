package sqlinject

import (
	"time"
)

type Sqlinject0114 struct{}

func NewSqlinject0114() *Sqlinject0114 {
	return &Sqlinject0114{}
}

func (e *Sqlinject0114) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0114) Name() string { return "Sqlinject0114" }
func (e *Sqlinject0114) Timestamp() time.Time { return time.Now() }
