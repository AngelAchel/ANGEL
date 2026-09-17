package orchestrator

import (
	"time"
)

// LayerConnector connects layers via event bus
type LayerConnector struct{}

func NewLayerConnector() *LayerConnector {
	return &LayerConnector{}
}

func (c *LayerConnector) Connect(source, dest string) ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "layerconnector:connected")
	return results, nil
}

func (c *LayerConnector) Name() string { return "LayerConnector" }
func (c *LayerConnector) Timestamp() time.Time { return time.Now() }
