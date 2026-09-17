package sqlinject

import (
	"time"
)

type Sqlinject0155 struct{}

func NewSqlinject0155() *Sqlinject0155 {
	return &Sqlinject0155{}
}

func (e *Sqlinject0155) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "sqlinject:done")
	return results, nil
}

func (e *Sqlinject0155) Name() string { return "Sqlinject0155" }
func (e *Sqlinject0155) Timestamp() time.Time { return time.Now() }
