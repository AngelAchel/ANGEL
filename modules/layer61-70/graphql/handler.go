package graphql

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
	eventbus.Publish(eventbus.NewEvent("event", "graphql", "*", "event", nil))
	eventbus.Publish(eventbus.NewEvent("event", "graphql", "*", "event", nil))
	return nil
}

func (h *Handler) Name() string         { return "Handler" }
func (h *Handler) Timestamp() time.Time { return time.Now() }
