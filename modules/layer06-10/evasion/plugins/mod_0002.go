package evasion

import (
	"time"
)

type Evasion0002 struct{}

func NewEvasion0002() *Evasion0002 {
	return &Evasion0002{}
}

func (e *Evasion0002) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "evasion:done")
	return results, nil
}

func (e *Evasion0002) Name() string { return "Evasion0002" }
func (e *Evasion0002) Timestamp() time.Time { return time.Now() }
