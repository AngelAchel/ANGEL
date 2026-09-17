package sqlinject

import (
	"time"
)

// InterLayerHandler handles inter-layer events
type InterLayerHandler struct{}

func NewInterLayerHandler() *InterLayerHandler {
	return &InterLayerHandler{}
}

func (h *InterLayerHandler) Handle(topic string) ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "interlayer:handled")
	return results, nil
}

func (h *InterLayerHandler) Name() string         { return "InterLayerHandler" }
func (h *InterLayerHandler) Timestamp() time.Time { return time.Now() }
