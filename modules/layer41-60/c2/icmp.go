package c2

import (
	"time"
)

type ICMP struct{}

func NewICMP() *ICMP {
	return &ICMP{}
}

func (e *ICMP) Ping(target string) ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "icmp:pinged")
	return results, nil
}

func (e *ICMP) Name() string         { return "ICMP" }
func (e *ICMP) Category() C2Category { return CategoryC2 }
func (e *ICMP) Timestamp() time.Time { return time.Now() }
