package c2server

import (
	"time"
)

type Ssl struct{}

func NewSsl() *Ssl {
	return &Ssl{}
}

func (e *Ssl) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "ssl:done")
	return results, nil
}

func (e *Ssl) Name() string { return "Ssl" }
func (e *Ssl) Timestamp() time.Time { return time.Now() }
