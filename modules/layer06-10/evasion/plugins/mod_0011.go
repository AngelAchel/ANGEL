package evasion

import (
	"time"
)

type Evasion0011 struct{}

func NewEvasion0011() *Evasion0011 {
	return &Evasion0011{}
}

func (e *Evasion0011) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "evasion:done")
	return results, nil
}

func (e *Evasion0011) Name() string { return "Evasion0011" }
func (e *Evasion0011) Timestamp() time.Time { return time.Now() }
