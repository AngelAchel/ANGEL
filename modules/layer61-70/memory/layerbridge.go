package memory

import (
	"time"
)

// LayerBridge bridges layers
type LayerBridge struct{}

func NewLayerBridge() *LayerBridge {
	return &LayerBridge{}
}

func (b *LayerBridge) Bridge() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "layerbridge:bridged")
	return results, nil
}

func (b *LayerBridge) Name() string { return "LayerBridge" }
func (b *LayerBridge) Timestamp() time.Time { return time.Now() }
