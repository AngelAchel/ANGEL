package evasion

import (
	"time"
)

type EvasionAgent0101 struct{}

func NewEvasionAgent0101() *EvasionAgent0101 {
	return &EvasionAgent0101{}
}

func (e *EvasionAgent0101) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "evasion:done")
	return results, nil
}

func (e *EvasionAgent0101) Name() string         { return "EvasionAgent0101" }
func (e *EvasionAgent0101) Timestamp() time.Time { return time.Now() }
