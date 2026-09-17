package orchestrator

import (
	"time"
)

// EventRouter routes events between layers
type EventRouter struct{}

func NewEventRouter() *EventRouter {
	return &EventRouter{}
}

func (r *EventRouter) Route(event Event) ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "eventrouter:routed")
	return results, nil
}

func (r *EventRouter) Name() string { return "EventRouter" }
func (r *EventRouter) Timestamp() time.Time { return time.Now() }
