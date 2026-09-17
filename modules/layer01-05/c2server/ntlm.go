package c2server

import (
	"time"
)

type Ntlm struct{}

func NewNtlm() *Ntlm {
	return &Ntlm{}
}

func (e *Ntlm) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "ntlm:done")
	return results, nil
}

func (e *Ntlm) Name() string         { return "Ntlm" }
func (e *Ntlm) Timestamp() time.Time { return time.Now() }
