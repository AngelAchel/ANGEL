package c2server

import (
	"time"
)

type Ssh struct{}

func NewSsh() *Ssh {
	return &Ssh{}
}

func (e *Ssh) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "ssh:done")
	return results, nil
}

func (e *Ssh) Name() string         { return "Ssh" }
func (e *Ssh) Timestamp() time.Time { return time.Now() }
