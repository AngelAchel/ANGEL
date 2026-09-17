package ledger

import (
	"sync"
	"time"
)

type Ledger struct {
	mu      sync.RWMutex
	entries []Entry
}

type Entry struct {
	ID        string    `json:"id"`
	Topic     string    `json:"topic"`
	Timestamp time.Time `json:"timestamp"`
	Source    string    `json:"source"`
	Data      string    `json:"data"`
	TraceID   string    `json:"trace_id"`
}

func NewLedger() *Ledger {
	return &Ledger{
		entries: make([]Entry, 0, 100),
	}
}

func (l *Ledger) Record(topic, source, data, traceID string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.entries = append(l.entries, Entry{
		ID:        generateID(),
		Topic:     topic,
		Timestamp: time.Now(),
		Source:    source,
		Data:      data,
		TraceID:   traceID,
	})
}

func (l *Ledger) Query(topic string) []Entry {
	l.mu.RLock()
	defer l.mu.RUnlock()
	var results []Entry
	for _, e := range l.entries {
		if e.Topic == topic {
			results = append(results, e)
		}
	}
	return results
}

func (l *Ledger) Name() string { return "Ledger" }
func (l *Ledger) Timestamp() time.Time { return time.Now() }
func (l *Ledger) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "ledger:running")
	return results, nil
}

func generateID() string {
	return "ledger-" + time.Now().Format("20060102150405")
}
