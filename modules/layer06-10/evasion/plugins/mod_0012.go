package evasion

import (
	"time"
)

type Evasion0012 struct{}

func NewEvasion0012() *Evasion0012 {
	return &Evasion0012{}
}

func (e *Evasion0012) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "evasion:done")
	return results, nil
}

func (e *Evasion0012) Name() string { return "Evasion0012" }
func (e *Evasion0012) Timestamp() time.Time { return time.Now() }
