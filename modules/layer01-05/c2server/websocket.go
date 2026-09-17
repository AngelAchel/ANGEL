package c2server

import (
	"time"
)

type Websocket struct{}

func NewWebsocket() *Websocket {
	return &Websocket{}
}

func (e *Websocket) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "websocket:done")
	return results, nil
}

func (e *Websocket) Name() string { return "Websocket" }
func (e *Websocket) Timestamp() time.Time { return time.Now() }
