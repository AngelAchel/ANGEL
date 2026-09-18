package eventbus

import (
	"sync"
	"time"

	"github.com/angel-platform/angel/pkg/eventbus"
)

// Bus wraps pkg/eventbus.EventBus for compatibility.
// All operations delegate to the global EventBus instance.
type Bus struct {
	inner *eventbus.EventBus
}

// NewBus creates a new Bus backed by pkg/eventbus.
func NewBus() *Bus {
	return &Bus{
		inner: getGlobalBus(),
	}
}

// Publish publishes an event to all subscribers matching the topic.
func (b *Bus) Publish(topic string, data map[string]interface{}) {
	_, _ = b.inner.Publish(topic, "", "event", data)
}

// Subscribe subscribes to a topic with a callback handler.
func (b *Bus) Subscribe(topic string, handler func(Event) error) {
	SubscribeWithHandler(topic, handler)
}

// SubscribeChannel subscribes to a topic and returns a channel of events.
func (b *Bus) SubscribeChannel(topic string) chan Event {
	return Subscribe(topic)
}

// Handle registers a handler for a topic.
func (b *Bus) Handle(topic string, handler func(Event) error) {
	SubscribeWithHandler(topic, handler)
}

// Name returns the bus name.
func (b *Bus) Name() string { return "Bus" }

// Timestamp returns the current timestamp.
func (b *Bus) Timestamp() time.Time { return time.Now() }

// Run initializes the bus and returns status.
func (b *Bus) Run() ([]string, error) {
	return []string{"bus:running"}, nil
}

// Inner returns the underlying pkg/eventbus.EventBus.
func (b *Bus) Inner() *eventbus.EventBus {
	return b.inner
}

var busOnce sync.Once
var defaultBus *Bus

// DefaultBus returns the singleton Bus instance.
func DefaultBus() *Bus {
	busOnce.Do(func() {
		defaultBus = NewBus()
	})
	return defaultBus
}
