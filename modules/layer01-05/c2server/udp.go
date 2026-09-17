package c2server

import (
	"time"
)

type Udp struct{}

func NewUdp() *Udp {
	return &Udp{}
}

func (e *Udp) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "udp:done")
	return results, nil
}

func (e *Udp) Name() string         { return "Udp" }
func (e *Udp) Timestamp() time.Time { return time.Now() }
