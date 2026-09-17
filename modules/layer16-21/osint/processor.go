package osint

import (
	"time"
)

// Processor processes events from other layers
type Processor struct{}

func NewProcessor() *Processor {
	return &Processor{}
}

func (p *Processor) Process(event Event) ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "processor:processed")
	return results, nil
}

func (p *Processor) Name() string         { return "Processor" }
func (p *Processor) Timestamp() time.Time { return time.Now() }

type Event struct {
	ID        string
	Topic     string
	Timestamp interface{}
	Source    string
	Dest      string
	Type      string
	Priority  int
	Data      map[string]interface{}
	TraceID   string
}
