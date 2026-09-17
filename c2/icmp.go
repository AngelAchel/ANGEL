package c2

import (
	"time"
)

type ICMP struct{}

func NewICMP() *ICMP {
	return &ICMP{}
}

func (i *ICMP) Ping() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "icmp:done")
	return results, nil
}

func (i *ICMP) Name() string { return "ICMP" }
func (i *ICMP) Timestamp() time.Time { return time.Now() }
