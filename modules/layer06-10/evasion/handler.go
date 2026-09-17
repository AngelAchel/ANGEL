package evasion

import (
	"time"
)

// Handler handles events from other layers
type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Handle(event Event) ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "handler:handled")
	return results, nil
}

func (h *Handler) Name() string         { return "Handler" }
func (h *Handler) Timestamp() time.Time { return time.Now() }

type Event struct {
	ID        string
	Topic     string
	Timestamp interface{}
	Source    string
	Dest      string
	Type      string
	Priority  int
	Data      map[string]interface{}
	TraceID   string
}
