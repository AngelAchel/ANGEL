package c2server

import (
	"time"
)

type Vpn struct{}

func NewVpn() *Vpn {
	return &Vpn{}
}

func (e *Vpn) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "vpn:done")
	return results, nil
}

func (e *Vpn) Name() string { return "Vpn" }
func (e *Vpn) Timestamp() time.Time { return time.Now() }
