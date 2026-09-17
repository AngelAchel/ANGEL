package c2server

import (
	"time"
)

type Gateway struct{}

func NewGateway() *Gateway {
	return &Gateway{}
}

func (e *Gateway) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "gateway:done")
	return results, nil
}

func (e *Gateway) Name() string         { return "Gateway" }
func (e *Gateway) Timestamp() time.Time { return time.Now() }
