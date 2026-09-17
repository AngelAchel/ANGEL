package ai

import (
	"time"
)

// Filter filters events between layers
type Filter struct{}

func NewFilter() *Filter {
	return &Filter{}
}

func (f *Filter) Filter(event Event) ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "filter:filtered")
	return results, nil
}

func (f *Filter) Name() string { return "Filter" }
func (f *Filter) Timestamp() time.Time { return time.Now() }

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
