package eventbus

import (
	"sync"
	"time"
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

var (
	subscribers = make(map[string][]chan Event)
	mu          sync.RWMutex
)

func Publish(event Event) {
	mu.RLock()
	defer mu.RUnlock()
	for topic, channels := range subscribers {
		if matchTopic(event.Topic, topic) {
			for _, ch := range channels {
				select {
				case ch <- event:
				default:
				}
			}
		}
	}
}

func Subscribe(topic string) chan Event {
	mu.Lock()
	defer mu.Unlock()
	ch := make(chan Event, 100)
	subscribers[topic] = append(subscribers[topic], ch)
	return ch
}

func matchTopic(eventTopic, subTopic string) bool {
	if subTopic == "*" {
		return true
	}
	return eventTopic == subTopic || len(eventTopic) >= len(subTopic)
}

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
	}
}

func generateID() string {
	return "evt-" + time.Now().Format("20060102150405")
}
