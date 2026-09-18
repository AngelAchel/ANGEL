package grpc

import (
	"time"

	"github.com/angel-platform/angel/modules/eventbus"
)

// Connector connects this layer to other layers
type Connector struct{}

func NewConnector() *Connector {
	return &Connector{}
}

func (c *Connector) Connect() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "grpc:connected")
	return results, nil
}

func (c *Connector) Name() string         { return "Connector" }
func (c *Connector) Timestamp() time.Time { return time.Now() }

// EventBusConnector connects via event bus
type EventBusConnector struct{}

func NewEventBusConnector() *EventBusConnector {
	return &EventBusConnector{}
}

func (e *EventBusConnector) Publish(topic string, data map[string]interface{}) error {
	event := eventbus.NewEvent(topic, "grpc", "*", "event", data)
	eventbus.Publish(event)
	return nil
}

func (e *EventBusConnector) Subscribe(topic string) chan eventbus.Event {
	return eventbus.Subscribe(topic)
}

func (e *EventBusConnector) Name() string         { return "EventBusConnector" }
func (e *EventBusConnector) Timestamp() time.Time { return time.Now() }
