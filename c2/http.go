package c2

import (
	"time"
)

type HTTP struct{}

func NewHTTP() *HTTP {
	return &HTTP{}
}

func (h *HTTP) Request() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "http:done")
	return results, nil
}

func (h *HTTP) Name() string { return "HTTP" }
func (h *HTTP) Timestamp() time.Time { return time.Now() }
