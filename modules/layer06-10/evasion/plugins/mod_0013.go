package evasion

import (
	"time"
)

type Evasion0013 struct{}

func NewEvasion0013() *Evasion0013 {
	return &Evasion0013{}
}

func (e *Evasion0013) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "evasion:done")
	return results, nil
}

func (e *Evasion0013) Name() string { return "Evasion0013" }
func (e *Evasion0013) Timestamp() time.Time { return time.Now() }
