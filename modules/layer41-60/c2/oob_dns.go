package c2

import (
	"time"
)

type OOBDNS struct{}

func NewOOBDNS() *OOBDNS {
	return &OOBDNS{}
}

func (o *OOBDNS) Resolve() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "oob_dns:done")
	return results, nil
}

func (o *OOBDNS) Name() string { return "OOBDNS" }
func (o *OOBDNS) Timestamp() time.Time { return time.Now() }
