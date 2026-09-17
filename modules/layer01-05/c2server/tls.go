package c2server

import (
	"time"
)

type Tls struct{}

func NewTls() *Tls {
	return &Tls{}
}

func (e *Tls) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "tls:done")
	return results, nil
}

func (e *Tls) Name() string { return "Tls" }
func (e *Tls) Timestamp() time.Time { return time.Now() }
