package c2

import (
	"time"
)

type DoH struct{}

func NewDoH() *DoH {
	return &DoH{}
}

func (d *DoH) Resolve() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "doh:done")
	return results, nil
}

func (d *DoH) Name() string { return "DoH" }
func (d *DoH) Timestamp() time.Time { return time.Now() }
