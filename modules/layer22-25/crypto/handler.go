package crypto

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
	evt := eventbus.NewEvent("layer:handle", "crypto", "*", "handle", map[string]interface{}{
		"event": event,
	})
	eventbus.Publish(evt)
	return nil
}

func (h *Handler) Name() string { return "Handler" }
func (h *Handler) Timestamp() time.Time { return time.Now() }
