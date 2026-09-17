package c2server

import (
	"time"
)

type Rpc struct{}

func NewRpc() *Rpc {
	return &Rpc{}
}

func (e *Rpc) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "rpc:done")
	return results, nil
}

func (e *Rpc) Name() string { return "Rpc" }
func (e *Rpc) Timestamp() time.Time { return time.Now() }
