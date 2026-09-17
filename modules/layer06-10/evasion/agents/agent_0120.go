package evasion

import (
	"time"
)

type EvasionAgent0120 struct{}

func NewEvasionAgent0120() *EvasionAgent0120 {
	return &EvasionAgent0120{}
}

func (e *EvasionAgent0120) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "evasion:done")
	return results, nil
}

func (e *EvasionAgent0120) Name() string { return "EvasionAgent0120" }
func (e *EvasionAgent0120) Timestamp() time.Time { return time.Now() }
