package evasion

import (
	"time"
)

type Evasion0000 struct{}

func NewEvasion0000() *Evasion0000 {
	return &Evasion0000{}
}

func (e *Evasion0000) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "evasion:done")
	return results, nil
}

func (e *Evasion0000) Name() string { return "Evasion0000" }
func (e *Evasion0000) Timestamp() time.Time { return time.Now() }
