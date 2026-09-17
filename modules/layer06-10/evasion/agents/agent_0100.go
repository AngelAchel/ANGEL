package evasion

import (
	"time"
)

type EvasionAgent0100 struct{}

func NewEvasionAgent0100() *EvasionAgent0100 {
	return &EvasionAgent0100{}
}

func (e *EvasionAgent0100) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "evasion:done")
	return results, nil
}

func (e *EvasionAgent0100) Name() string { return "EvasionAgent0100" }
func (e *EvasionAgent0100) Timestamp() time.Time { return time.Now() }
