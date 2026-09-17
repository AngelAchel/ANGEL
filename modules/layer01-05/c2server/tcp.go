package c2server

import (
	"time"
)

type Tcp struct{}

func NewTcp() *Tcp {
	return &Tcp{}
}

func (e *Tcp) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "tcp:done")
	return results, nil
}

func (e *Tcp) Name() string { return "Tcp" }
func (e *Tcp) Timestamp() time.Time { return time.Now() }
