package sqlinject

import (
	"time"
)

type Sqlinject0166 struct{}

func NewSqlinject0166() *Sqlinject0166 {
	return &Sqlinject0166{}
}

func (e *Sqlinject0166) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0166) Name() string { return "Sqlinject0166" }
func (e *Sqlinject0166) Timestamp() time.Time { return time.Now() }
