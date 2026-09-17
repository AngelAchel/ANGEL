package storage

import (
	"testing"
)

func TestAgentsList(t *testing.T) {
	a := NewAgents()
	if a.Name() != "Agents" {
		t.Errorf("expected Agents, got %s", a.Name())
	}
}
