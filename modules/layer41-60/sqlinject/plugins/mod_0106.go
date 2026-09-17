package sqlinject

import (
	"time"
)

type Sqlinject0106 struct{}

func NewSqlinject0106() *Sqlinject0106 {
	return &Sqlinject0106{}
}

func (e *Sqlinject0106) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0106) Name() string { return "Sqlinject0106" }
func (e *Sqlinject0106) Timestamp() time.Time { return time.Now() }
