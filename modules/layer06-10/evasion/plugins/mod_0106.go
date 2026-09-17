package evasion

import (
	"time"
)

type Evasion0106 struct{}

func NewEvasion0106() *Evasion0106 {
	return &Evasion0106{}
}

func (e *Evasion0106) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "evasion:done")
	return results, nil
}

func (e *Evasion0106) Name() string { return "Evasion0106" }
func (e *Evasion0106) Timestamp() time.Time { return time.Now() }
