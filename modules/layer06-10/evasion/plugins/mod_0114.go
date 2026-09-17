package evasion

import (
	"time"
)

type Evasion0114 struct{}

func NewEvasion0114() *Evasion0114 {
	return &Evasion0114{}
}

func (e *Evasion0114) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "evasion:done")
	return results, nil
}

func (e *Evasion0114) Name() string { return "Evasion0114" }
func (e *Evasion0114) Timestamp() time.Time { return time.Now() }
