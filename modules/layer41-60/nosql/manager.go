package nosql

import (
	"time"
)

// Manager manages cross-layer operations
type Manager struct{}

func NewManager() *Manager {
	return &Manager{}
}

func (m *Manager) Manage() error {
	return nil
}

func (m *Manager) Name() string { return "Manager" }
func (m *Manager) Timestamp() time.Time { return time.Now() }
