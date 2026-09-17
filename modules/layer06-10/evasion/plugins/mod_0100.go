package evasion

import (
	"time"
)

type Evasion0100 struct{}

func NewEvasion0100() *Evasion0100 {
	return &Evasion0100{}
}

func (e *Evasion0100) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "evasion:done")
	return results, nil
}

func (e *Evasion0100) Name() string { return "Evasion0100" }
func (e *Evasion0100) Timestamp() time.Time { return time.Now() }
