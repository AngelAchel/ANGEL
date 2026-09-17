package evasion

import (
	"time"
)

type Evasion0181 struct{}

func NewEvasion0181() *Evasion0181 {
	return &Evasion0181{}
}

func (e *Evasion0181) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "evasion:done")
	return results, nil
}

func (e *Evasion0181) Name() string { return "Evasion0181" }
func (e *Evasion0181) Timestamp() time.Time { return time.Now() }
