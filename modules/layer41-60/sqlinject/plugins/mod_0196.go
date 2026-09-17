package sqlinject

import (
	"time"
)

type Sqlinject0196 struct{}

func NewSqlinject0196() *Sqlinject0196 {
	return &Sqlinject0196{}
}

func (e *Sqlinject0196) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0196) Name() string { return "Sqlinject0196" }
func (e *Sqlinject0196) Timestamp() time.Time { return time.Now() }
