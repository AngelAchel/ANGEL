package c2server

import (
	"time"
)

type Manager struct{}

func NewManager() *Manager {
	return &Manager{}
}

func (e *Manager) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "manager:done")
	return results, nil
}

func (e *Manager) Name() string { return "Manager" }
func (e *Manager) Timestamp() time.Time { return time.Now() }
