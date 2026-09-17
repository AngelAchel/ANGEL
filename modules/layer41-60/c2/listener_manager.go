package c2

import (
	"time"
)

type ListenerManager struct{}

func NewListenerManager() *ListenerManager {
	return &ListenerManager{}
}

func (l *ListenerManager) Manage() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "listener_manager:done")
	return results, nil
}

func (l *ListenerManager) Name() string { return "ListenerManager" }
func (l *ListenerManager) Timestamp() time.Time { return time.Now() }
