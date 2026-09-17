package c2

import (
	"time"
)

type DNS struct{}

func NewDNS() *DNS {
	return &DNS{}
}

func (d *DNS) Query() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "dns:done")
	return results, nil
}

func (d *DNS) Name() string { return "DNS" }
func (d *DNS) Timestamp() time.Time { return time.Now() }
