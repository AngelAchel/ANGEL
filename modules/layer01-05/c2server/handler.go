package c2server

import (
	"time"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (e *Handler) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "handler:done")
	return results, nil
}

func (e *Handler) Name() string         { return "Handler" }
func (e *Handler) Timestamp() time.Time { return time.Now() }
