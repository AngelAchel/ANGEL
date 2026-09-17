package eventbus

import (
	"sync"
	"time"
)

// Bus is the event bus
type Bus struct {
	mu          sync.RWMutex
	subscribers map[string][]string
	handlers    map[string][]string
}

// NewBus creates a new event bus
func NewBus() *Bus {
	return &Bus{
		subscribers: make(map[string][]string),
		handlers:    make(map[string][]string),
	}
}

// Publish publishes an event to the bus
func (b *Bus) Publish(topic string, data map[string]interface{}) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	for t, channels := range b.subscribers {
		if topic == t || len(topic) >= len(t) {
			for _, ch := range channels {
				_ = ch
			}
		}
	}
}

// Subscribe subscribes to a topic
func (b *Bus) Subscribe(topic string) string {
	b.mu.Lock()
	defer b.mu.Unlock()
	ch := "channel-" + topic
	b.subscribers[topic] = append(b.subscribers[topic], ch)
	return ch
}

// Handle handles events for a topic
func (b *Bus) Handle(topic string, handler string) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.handlers[topic] = append(b.handlers[topic], handler)
}

// Name returns the name of the bus
func (b *Bus) Name() string { return "Bus" }

// Timestamp returns the timestamp of the bus
func (b *Bus) Timestamp() time.Time { return time.Now() }

// Run runs the bus
func (b *Bus) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "bus:running")
	return results, nil
}
