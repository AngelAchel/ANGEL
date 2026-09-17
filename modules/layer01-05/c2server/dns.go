package c2server

import (
	"time"
)

type Dns struct{}

func NewDns() *Dns {
	return &Dns{}
}

func (e *Dns) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "dns:done")
	return results, nil
}

func (e *Dns) Name() string         { return "Dns" }
func (e *Dns) Timestamp() time.Time { return time.Now() }
