package c2server

import (
	"time"
)

type Grpc struct{}

func NewGrpc() *Grpc {
	return &Grpc{}
}

func (e *Grpc) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "grpc:done")
	return results, nil
}

func (e *Grpc) Name() string { return "Grpc" }
func (e *Grpc) Timestamp() time.Time { return time.Now() }
