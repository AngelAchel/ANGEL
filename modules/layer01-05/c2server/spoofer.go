package c2server

import (
	"time"
)

type Spoofer struct{}

func NewSpoofer() *Spoofer {
	return &Spoofer{}
}

func (e *Spoofer) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "spoofer:done")
	return results, nil
}

func (e *Spoofer) Name() string { return "Spoofer" }
func (e *Spoofer) Timestamp() time.Time { return time.Now() }
