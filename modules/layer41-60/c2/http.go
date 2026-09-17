package c2

import (
	"time"
)

type HTTP struct{}

func NewHTTP() *HTTP {
	return &HTTP{}
}

func (e *HTTP) Request(url string) ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "http:requested")
	return results, nil
}

func (e *HTTP) Name() string { return "HTTP" }
func (e *HTTP) Category() C2Category { return CategoryC2 }
func (e *HTTP) Timestamp() time.Time { return time.Now() }
