package crypto

import (
	"time"
)

// Aggregator aggregates events from all layers
type Aggregator struct{}

func NewAggregator() *Aggregator {
	return &Aggregator{}
}

func (a *Aggregator) Aggregate(events []Event) ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "aggregator:aggregated")
	return results, nil
}

func (a *Aggregator) Name() string { return "Aggregator" }
func (a *Aggregator) Timestamp() time.Time { return time.Now() }

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
