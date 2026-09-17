package evasion

import (
	"time"
)

type Evasion0083 struct{}

func NewEvasion0083() *Evasion0083 {
	return &Evasion0083{}
}

func (e *Evasion0083) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "evasion:done")
	return results, nil
}

func (e *Evasion0083) Name() string { return "Evasion0083" }
func (e *Evasion0083) Timestamp() time.Time { return time.Now() }
