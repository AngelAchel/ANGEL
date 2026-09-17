package evasion

import (
	"time"
)

type Evasion0037 struct{}

func NewEvasion0037() *Evasion0037 {
	return &Evasion0037{}
}

func (e *Evasion0037) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "evasion:done")
	return results, nil
}

func (e *Evasion0037) Name() string { return "Evasion0037" }
func (e *Evasion0037) Timestamp() time.Time { return time.Now() }
