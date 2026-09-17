package evasion

import (
	"time"
)

type Evasion0005 struct{}

func NewEvasion0005() *Evasion0005 {
	return &Evasion0005{}
}

func (e *Evasion0005) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "evasion:done")
	return results, nil
}

func (e *Evasion0005) Name() string { return "Evasion0005" }
func (e *Evasion0005) Timestamp() time.Time { return time.Now() }
