package credential

import (
	"time"
)

// Router routes events between layers
type Router struct{}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Route(event Event) ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "router:routed")
	return results, nil
}

func (r *Router) Name() string { return "Router" }
func (r *Router) Timestamp() time.Time { return time.Now() }

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
