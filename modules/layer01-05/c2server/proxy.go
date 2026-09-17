package c2server

import (
	"time"
)

type Proxy struct{}

func NewProxy() *Proxy {
	return &Proxy{}
}

func (e *Proxy) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "proxy:done")
	return results, nil
}

func (e *Proxy) Name() string         { return "Proxy" }
func (e *Proxy) Timestamp() time.Time { return time.Now() }
