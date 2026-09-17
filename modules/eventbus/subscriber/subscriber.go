package subscriber

import (
	"time"
)

type Subscriber struct {
	topics   map[string][]chan Event
	handlers map[string][]func(Event)
}

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

func NewSubscriber() *Subscriber {
	return &Subscriber{
		topics:   make(map[string][]chan Event),
		handlers: make(map[string][]func(Event)),
	}
}

func (s *Subscriber) Subscribe(topic string) chan Event {
	ch := make(chan Event, 100)
	s.topics[topic] = append(s.topics[topic], ch)
	return ch
}

func (s *Subscriber) Handle(topic string, handler func(Event)) {
	s.handlers[topic] = append(s.handlers[topic], handler)
}

func (s *Subscriber) Name() string { return "Subscriber" }
func (s *Subscriber) Timestamp() time.Time { return time.Now() }
func (s *Subscriber) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "subscriber:running")
	return results, nil
}
