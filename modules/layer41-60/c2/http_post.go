package c2

import (
	"time"
)

type HTTPPost struct{}

func NewHTTPPost() *HTTPPost {
	return &HTTPPost{}
}

func (h *HTTPPost) Post() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "http_post:done")
	return results, nil
}

func (h *HTTPPost) Name() string         { return "HTTPPost" }
func (h *HTTPPost) Timestamp() time.Time { return time.Now() }
