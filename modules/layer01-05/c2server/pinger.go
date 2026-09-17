package c2server

import (
	"time"
)

type Pinger struct{}

func NewPinger() *Pinger {
	return &Pinger{}
}

func (e *Pinger) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "pinger:done")
	return results, nil
}

func (e *Pinger) Name() string { return "Pinger" }
func (e *Pinger) Timestamp() time.Time { return time.Now() }
