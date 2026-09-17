package c2server

import (
	"time"
)

type Icmp struct{}

func NewIcmp() *Icmp {
	return &Icmp{}
}

func (e *Icmp) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "icmp:done")
	return results, nil
}

func (e *Icmp) Name() string { return "Icmp" }
func (e *Icmp) Timestamp() time.Time { return time.Now() }
