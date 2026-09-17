package evasion

import (
	"time"
)

type Evasion0073 struct{}

func NewEvasion0073() *Evasion0073 {
	return &Evasion0073{}
}

func (e *Evasion0073) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "evasion:done")
	return results, nil
}

func (e *Evasion0073) Name() string { return "Evasion0073" }
func (e *Evasion0073) Timestamp() time.Time { return time.Now() }
