package c2server

import (
	"time"
)

type Tunnel struct{}

func NewTunnel() *Tunnel {
	return &Tunnel{}
}

func (e *Tunnel) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "tunnel:done")
	return results, nil
}

func (e *Tunnel) Name() string { return "Tunnel" }
func (e *Tunnel) Timestamp() time.Time { return time.Now() }
