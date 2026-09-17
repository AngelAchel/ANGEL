package orchestrator

import (
	"sync"
	"time"
)

// Orchestrator manages all layers
type Orchestrator struct {
	mu       sync.RWMutex
	layers   map[string]Layer
	handlers map[string]func(Event)
}

type Layer struct {
	Name    string
	Version string
	Active  bool
}

type Event struct {
	ID        string
	Topic     string
	Timestamp time.Time
	Source    string
	Dest      string
	Type      string
	Priority  int
	Data      map[string]interface{}
	TraceID   string
}

func NewOrchestrator() *Orchestrator {
	return &Orchestrator{
		layers:   make(map[string]Layer),
		handlers: make(map[string]func(Event)),
	}
}

func (o *Orchestrator) RegisterLayer(name, version string) {
	o.mu.Lock()
	defer o.mu.Unlock()
	o.layers[name] = Layer{Name: name, Version: version, Active: true}
}

func (o *Orchestrator) Publish(topic string, event Event) {
	o.mu.RLock()
	defer o.mu.RUnlock()
	for _, handler := range o.handlers {
		handler(event)
	}
}

func (o *Orchestrator) Subscribe(topic string, handler func(Event)) {
	o.mu.Lock()
	defer o.mu.Unlock()
	o.handlers[topic] = handler
}

func (o *Orchestrator) Name() string { return "Orchestrator" }
func (o *Orchestrator) Timestamp() time.Time { return time.Now() }
func (o *Orchestrator) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "orchestrator:running")
	return results, nil
}
