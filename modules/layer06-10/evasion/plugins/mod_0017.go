package evasion

import (
	"time"
)

type Evasion0017 struct{}

func NewEvasion0017() *Evasion0017 {
	return &Evasion0017{}
}

func (e *Evasion0017) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "evasion:done")
	return results, nil
}

func (e *Evasion0017) Name() string { return "Evasion0017" }
func (e *Evasion0017) Timestamp() time.Time { return time.Now() }
