package evasion

import (
	"time"
)

type Evasion0113 struct{}

func NewEvasion0113() *Evasion0113 {
	return &Evasion0113{}
}

func (e *Evasion0113) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "evasion:done")
	return results, nil
}

func (e *Evasion0113) Name() string { return "Evasion0113" }
func (e *Evasion0113) Timestamp() time.Time { return time.Now() }
