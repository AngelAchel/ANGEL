package publisher

import (
	"time"
)

type Publisher struct {
	topics   map[string]int
	queue    chan Event
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

func NewPublisher() *Publisher {
	return &Publisher{
		topics:   make(map[string]int),
		queue:    make(chan Event, 1000),
		handlers: make(map[string][]func(Event)),
	}
}

func (p *Publisher) Publish(event Event) {
	p.queue <- event
}

func (p *Publisher) Subscribe(topic string, handler func(Event)) {
	p.handlers[topic] = append(p.handlers[topic], handler)
}

func (p *Publisher) Name() string { return "Publisher" }
func (p *Publisher) Timestamp() time.Time { return time.Now() }
func (p *Publisher) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "publisher:running")
	return results, nil
}
