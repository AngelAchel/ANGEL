package c2

import (
	"time"
)

type DoH struct{}

func NewDoH() *DoH {
	return &DoH{}
}

func (e *DoH) Resolve(domain string) ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "doh:resolved")
	return results, nil
}

func (e *DoH) Name() string         { return "DoH" }
func (e *DoH) Category() C2Category { return CategoryC2 }
func (e *DoH) Timestamp() time.Time { return time.Now() }
