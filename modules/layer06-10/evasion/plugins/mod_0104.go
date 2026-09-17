package evasion

import (
	"time"
)

type Evasion0104 struct{}

func NewEvasion0104() *Evasion0104 {
	return &Evasion0104{}
}

func (e *Evasion0104) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "evasion:done")
	return results, nil
}

func (e *Evasion0104) Name() string { return "Evasion0104" }
func (e *Evasion0104) Timestamp() time.Time { return time.Now() }
