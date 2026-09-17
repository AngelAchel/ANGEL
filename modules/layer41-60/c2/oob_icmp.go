package c2

import (
	"time"
)

type OOBICMP struct{}

func NewOOBICMP() *OOBICMP {
	return &OOBICMP{}
}

func (o *OOBICMP) Ping() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "oob_icmp:done")
	return results, nil
}

func (o *OOBICMP) Name() string { return "OOBICMP" }
func (o *OOBICMP) Timestamp() time.Time { return time.Now() }
