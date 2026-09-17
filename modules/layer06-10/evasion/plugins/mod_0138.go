package evasion

import (
	"time"
)

type Evasion0138 struct{}

func NewEvasion0138() *Evasion0138 {
	return &Evasion0138{}
}

func (e *Evasion0138) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "evasion:done")
	return results, nil
}

func (e *Evasion0138) Name() string { return "Evasion0138" }
func (e *Evasion0138) Timestamp() time.Time { return time.Now() }
