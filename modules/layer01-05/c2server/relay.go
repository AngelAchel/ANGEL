package c2server

import (
	"time"
)

type Relay struct{}

func NewRelay() *Relay {
	return &Relay{}
}

func (e *Relay) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "relay:done")
	return results, nil
}

func (e *Relay) Name() string { return "Relay" }
func (e *Relay) Timestamp() time.Time { return time.Now() }
