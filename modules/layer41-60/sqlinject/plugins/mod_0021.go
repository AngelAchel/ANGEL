package sqlinject

import (
	"time"
)

type Sqlinject0021 struct{}

func NewSqlinject0021() *Sqlinject0021 {
	return &Sqlinject0021{}
}

func (e *Sqlinject0021) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0021) Name() string { return "Sqlinject0021" }
func (e *Sqlinject0021) Timestamp() time.Time { return time.Now() }
