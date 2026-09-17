package network

import (
	"time"
)

// Mapper maps events between layers
type Mapper struct{}

func NewMapper() *Mapper {
	return &Mapper{}
}

func (m *Mapper) Map(event Event) ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "mapper:mapped")
	return results, nil
}

func (m *Mapper) Name() string         { return "Mapper" }
func (m *Mapper) Timestamp() time.Time { return time.Now() }

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
