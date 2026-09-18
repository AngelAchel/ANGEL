package eventbus

import (
	"sync"
	"time"

	"github.com/angel-platform/angel/pkg/eventbus"
)

// Event represents a message on the event bus
type Event struct {
	ID        string                 `json:"id"`
	Topic     string                 `json:"topic"`
	Timestamp time.Time              `json:"timestamp"`
	Source    string                 `json:"source"`
	Dest      string                 `json:"dest"`
	Type      string                 `json:"type"`
	Priority  int                    `json:"priority"`
	Data      map[string]interface{} `json:"data"`
	TraceID   string                 `json:"trace_id"`
}

// globalBus wraps pkg/eventbus.EventBus with HMAC signing and event log
var (
	globalBus *eventbus.EventBus
	globalMu  sync.Once
)

func getGlobalBus() *eventbus.EventBus {
	globalMu.Do(func() {
		globalBus = eventbus.New("angel-module-bridge-key")
	})
	return globalBus
}

// Publish publishes an event to all subscribers matching the topic.
// Delegates to pkg/eventbus.EventBus with HMAC signing and event log.
func Publish(event Event) {
	bus := getGlobalBus()
	data := map[string]interface{}{
		"data": event.Data,
	}
	_, _ = bus.Publish(event.Topic, event.Source, event.Type, data)
}

// Subscribe subscribes to a topic and returns a channel
// that receives matching events. Supports wildcard "layer:*" patterns.
func Subscribe(topic string) chan Event {
	ch := make(chan Event, 100)
	bus := getGlobalBus()

	handler := func(e eventbus.Event) error {
		ev := Event{
			ID:        e.ID,
			Topic:     e.Topic,
			Timestamp: time.Now(),
			Source:    e.Source,
			Dest:      e.Dest,
			Type:      e.Type,
			Priority:  e.Priority,
			Data:      e.Data,
			TraceID:   e.TraceID,
		}
		select {
		case ch <- ev:
		default:
		}
		return nil
	}

	bus.Subscribe(topic, handler)
	return ch
}

// NewEvent creates a new event with the given parameters.
func NewEvent(topic, source, dest, eventType string, data map[string]interface{}) Event {
	return Event{
		ID:        generateID(),
		Topic:     topic,
		Timestamp: time.Now(),
		Source:    source,
		Dest:      dest,
		Type:      eventType,
		Priority:  50,
		Data:      data,
		TraceID:   generateID(),
	}
}

// MatchTopic checks if eventTopic matches the subscription pattern.
// Supports exact match, wildcard "*", and prefix match for layer paths.
func MatchTopic(eventTopic, subTopic string) bool {
	if subTopic == "*" {
		return true
	}
	if subTopic == eventTopic {
		return true
	}
	// Prefix match: "layer01-05" matches "layer01-05.c2" but not "layer01-052"
	if len(eventTopic) > len(subTopic) && eventTopic[len(subTopic)] == '.' {
		return eventTopic[:len(subTopic)] == subTopic
	}
	return false
}

// GetGlobalBus returns the underlying pkg/eventbus.EventBus for direct use
// when more control is needed (e.g., in orchestrator or gateway).
func GetGlobalBus() *eventbus.EventBus {
	return getGlobalBus()
}

// SubscribeWithHandler subscribes to a topic with a callback function.
func SubscribeWithHandler(topic string, handler func(Event) error) {
	bus := getGlobalBus()
	wrapped := func(e eventbus.Event) error {
		ev := Event{
			ID:        e.ID,
			Topic:     e.Topic,
			Timestamp: time.Now(),
			Source:    e.Source,
			Dest:      e.Dest,
			Type:      e.Type,
			Priority:  e.Priority,
			Data:      e.Data,
			TraceID:   e.TraceID,
		}
		return handler(ev)
	}
	bus.Subscribe(topic, wrapped)
}

// Stop shuts down the global eventbus (for cleanup).
func Stop() {
	if globalBus != nil {
		globalBus.Stop()
	}
}

func generateID() string {
	return time.Now().UTC().Format("20060102150405")
}
