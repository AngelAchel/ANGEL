package sqlinject

import (
	"time"
)

// Transformer transforms events between layers
type Transformer struct{}

func NewTransformer() *Transformer {
	return &Transformer{}
}

func (t *Transformer) Transform(event Event) ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "transformer:transformed")
	return results, nil
}

func (t *Transformer) Name() string         { return "Transformer" }
func (t *Transformer) Timestamp() time.Time { return time.Now() }

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
