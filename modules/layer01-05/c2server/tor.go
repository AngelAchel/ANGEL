package c2server

import (
	"time"
)

type Tor struct{}

func NewTor() *Tor {
	return &Tor{}
}

func (e *Tor) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "tor:done")
	return results, nil
}

func (e *Tor) Name() string         { return "Tor" }
func (e *Tor) Timestamp() time.Time { return time.Now() }
