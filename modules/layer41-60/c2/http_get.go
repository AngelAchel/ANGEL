package c2

import (
	"time"
)

type HTTPGet struct{}

func NewHTTPGet() *HTTPGet {
	return &HTTPGet{}
}

func (h *HTTPGet) Get() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "http_get:done")
	return results, nil
}

func (h *HTTPGet) Name() string { return "HTTPGet" }
func (h *HTTPGet) Timestamp() time.Time { return time.Now() }
