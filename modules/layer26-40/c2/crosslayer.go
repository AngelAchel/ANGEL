package c2

import (
	"time"
)

// CrossLayerHandler handles cross-layer events
type CrossLayerHandler struct{}

func NewCrossLayerHandler() *CrossLayerHandler {
	return &CrossLayerHandler{}
}

func (h *CrossLayerHandler) Handle(topic string) ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "crosslayer:handled")
	return results, nil
}

func (h *CrossLayerHandler) Name() string { return "CrossLayerHandler" }
func (h *CrossLayerHandler) Timestamp() time.Time { return time.Now() }
