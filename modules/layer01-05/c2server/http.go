package c2server

import (
	"time"
)

type Http struct{}

func NewHttp() *Http {
	return &Http{}
}

func (e *Http) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "http:done")
	return results, nil
}

func (e *Http) Name() string { return "Http" }
func (e *Http) Timestamp() time.Time { return time.Now() }
