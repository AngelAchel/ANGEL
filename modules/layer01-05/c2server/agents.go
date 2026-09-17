package c2server

import (
	"time"
)

type Agents struct{}

func NewAgents() *Agents {
	return &Agents{}
}

func (a *Agents) List() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "agents:listed")
	return results, nil
}

func (a *Agents) Name() string { return "Agents" }
func (a *Agents) Timestamp() time.Time { return time.Now() }
