package evasion

import (
	"time"
)

type Evasion0118 struct{}

func NewEvasion0118() *Evasion0118 {
	return &Evasion0118{}
}

func (e *Evasion0118) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "evasion:done")
	return results, nil
}

func (e *Evasion0118) Name() string { return "Evasion0118" }
func (e *Evasion0118) Timestamp() time.Time { return time.Now() }
