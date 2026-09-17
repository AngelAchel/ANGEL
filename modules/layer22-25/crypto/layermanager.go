package crypto

import (
	"time"
)

// LayerManager manages layer connections
type LayerManager struct{}

func NewLayerManager() *LayerManager {
	return &LayerManager{}
}

func (m *LayerManager) Connect() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "layermanager:connected")
	return results, nil
}

func (m *LayerManager) Name() string         { return "LayerManager" }
func (m *LayerManager) Timestamp() time.Time { return time.Now() }
