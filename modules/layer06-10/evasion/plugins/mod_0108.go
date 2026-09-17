package evasion

import (
	"time"
)

type Evasion0108 struct{}

func NewEvasion0108() *Evasion0108 {
	return &Evasion0108{}
}

func (e *Evasion0108) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "evasion:done")
	return results, nil
}

func (e *Evasion0108) Name() string { return "Evasion0108" }
func (e *Evasion0108) Timestamp() time.Time { return time.Now() }
