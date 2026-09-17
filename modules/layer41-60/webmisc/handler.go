package webmisc

import (
	"time"
)

// Handler handles cross-layer events
type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Handle(event string) error {
	return nil
}

func (h *Handler) Name() string { return "Handler" }
func (h *Handler) Timestamp() time.Time { return time.Now() }
