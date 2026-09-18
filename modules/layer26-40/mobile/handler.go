package mobile

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
	eventbus.Publish(eventbus.NewEvent("event", "mobile", "*", "event", nil))
	eventbus.Publish(eventbus.NewEvent("event", "mobile", "*", "event", nil))
	return nil
}

func (h *Handler) Name() string         { return "Handler" }
func (h *Handler) Timestamp() time.Time { return time.Now() }
