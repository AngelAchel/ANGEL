package evasion

import (
	"time"
)

type Evasion0128 struct{}

func NewEvasion0128() *Evasion0128 {
	return &Evasion0128{}
}

func (e *Evasion0128) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "evasion:done")
	return results, nil
}

func (e *Evasion0128) Name() string { return "Evasion0128" }
func (e *Evasion0128) Timestamp() time.Time { return time.Now() }
