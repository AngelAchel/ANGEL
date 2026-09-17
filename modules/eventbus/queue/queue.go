package queue

import (
	"sync"
	"time"
)

type Queue struct {
	mu      sync.RWMutex
	items   []Event
	maxSize int
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

func NewQueue(maxSize int) *Queue {
	return &Queue{
		items:   make([]Event, 0, maxSize),
		maxSize: maxSize,
	}
}

func (q *Queue) Push(event Event) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.items) < q.maxSize {
		q.items = append(q.items, event)
	}
}

func (q *Queue) Pop() (Event, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.items) == 0 {
		return Event{}, false
	}
	event := q.items[0]
	q.items = q.items[1:]
	return event, true
}

func (q *Queue) Name() string         { return "Queue" }
func (q *Queue) Timestamp() time.Time { return time.Now() }
func (q *Queue) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "queue:running")
	return results, nil
}
