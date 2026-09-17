package c2server

import (
	"time"
)

type Kerberos struct{}

func NewKerberos() *Kerberos {
	return &Kerberos{}
}

func (e *Kerberos) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "kerberos:done")
	return results, nil
}

func (e *Kerberos) Name() string         { return "Kerberos" }
func (e *Kerberos) Timestamp() time.Time { return time.Now() }
