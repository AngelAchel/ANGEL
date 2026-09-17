package c2server

import (
	"time"
)

type Https struct{}

func NewHttps() *Https {
	return &Https{}
}

func (e *Https) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "https:done")
	return results, nil
}

func (e *Https) Name() string         { return "Https" }
func (e *Https) Timestamp() time.Time { return time.Now() }
