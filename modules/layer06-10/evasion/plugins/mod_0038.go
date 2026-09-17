package evasion

import (
	"time"
)

type Evasion0038 struct{}

func NewEvasion0038() *Evasion0038 {
	return &Evasion0038{}
}

func (e *Evasion0038) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "evasion:done")
	return results, nil
}

func (e *Evasion0038) Name() string { return "Evasion0038" }
func (e *Evasion0038) Timestamp() time.Time { return time.Now() }
