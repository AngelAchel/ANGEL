package storage

import (
	"time"
)

type Handlers struct{}

func NewHandlers() *Handlers {
	return &Handlers{}
}

func (h *Handlers) Handle() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "handlers:done")
	return results, nil
}

func (h *Handlers) Name() string         { return "Handlers" }
func (h *Handlers) Timestamp() time.Time { return time.Now() }
