package evasion

import (
	"time"
)

type Evasion0001 struct{}

func NewEvasion0001() *Evasion0001 {
	return &Evasion0001{}
}

func (e *Evasion0001) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "evasion:done")
	return results, nil
}

func (e *Evasion0001) Name() string { return "Evasion0001" }
func (e *Evasion0001) Timestamp() time.Time { return time.Now() }
