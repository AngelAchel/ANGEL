package evasion

import (
	"time"
)

type EvasionAgent0000 struct{}

func NewEvasionAgent0000() *EvasionAgent0000 {
	return &EvasionAgent0000{}
}

func (e *EvasionAgent0000) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "evasion:done")
	return results, nil
}

func (e *EvasionAgent0000) Name() string { return "EvasionAgent0000" }
func (e *EvasionAgent0000) Timestamp() time.Time { return time.Now() }
