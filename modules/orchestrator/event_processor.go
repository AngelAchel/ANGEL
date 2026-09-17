package orchestrator

import (
	"time"
)

// EventProcessor processes events from all layers
type EventProcessor struct{}

func NewEventProcessor() *EventProcessor {
	return &EventProcessor{}
}

func (p *EventProcessor) Process(event Event) ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "eventprocessor:processed")
	return results, nil
}

func (p *EventProcessor) Name() string { return "EventProcessor" }
func (p *EventProcessor) Timestamp() time.Time { return time.Now() }
