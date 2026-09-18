package ai

import (
	"time"

	"github.com/angel-platform/angel/modules/eventbus"
)

// Handler handles cross-layer events
type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Handle(event string) error {
	// Publish handler event
	evt := eventbus.NewEvent("layer:handle", "ai", "*", "handle", map[string]interface{}{
		"event": event,
	})
	eventbus.Publish(evt)
	eventbus.Publish(eventbus.NewEvent("event", "ai", "*", "event", nil))
	eventbus.Publish(eventbus.NewEvent("event", "ai", "*", "event", nil))
	return nil
}

func (h *Handler) Name() string         { return "Handler" }
func (h *Handler) Timestamp() time.Time { return time.Now() }
