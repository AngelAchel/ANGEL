package evasion

import (
	"time"
)

type Evasion0088 struct{}

func NewEvasion0088() *Evasion0088 {
	return &Evasion0088{}
}

func (e *Evasion0088) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "evasion:done")
	return results, nil
}

func (e *Evasion0088) Name() string { return "Evasion0088" }
func (e *Evasion0088) Timestamp() time.Time { return time.Now() }
