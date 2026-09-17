package c2

import (
	"time"
)

type DNS struct{}

func NewDNS() *DNS {
	return &DNS{}
}

func (e *DNS) Query(domain string) ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "dns:queried")
	return results, nil
}

func (e *DNS) Name() string         { return "DNS" }
func (e *DNS) Category() C2Category { return CategoryC2 }
func (e *DNS) Timestamp() time.Time { return time.Now() }
