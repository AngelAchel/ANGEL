package c2server

import (
	"time"
)

type Probe struct{}

func NewProbe() *Probe {
	return &Probe{}
}

func (e *Probe) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "probe:done")
	return results, nil
}

func (e *Probe) Name() string         { return "Probe" }
func (e *Probe) Timestamp() time.Time { return time.Now() }
