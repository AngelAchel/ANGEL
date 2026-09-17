package sqlinject

import (
	"time"
)

type Sqlinject0095 struct{}

func NewSqlinject0095() *Sqlinject0095 {
	return &Sqlinject0095{}
}

func (e *Sqlinject0095) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0095) Name() string { return "Sqlinject0095" }
func (e *Sqlinject0095) Timestamp() time.Time { return time.Now() }
