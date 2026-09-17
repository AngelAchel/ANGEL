package credential

import (
	"time"
)

type CredentialAgent0108 struct{}

func NewCredentialAgent0108() *CredentialAgent0108 {
	return &CredentialAgent0108{}
}

func (e *CredentialAgent0108) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "credential:done")
	return results, nil
}

func (e *CredentialAgent0108) Name() string         { return "CredentialAgent0108" }
func (e *CredentialAgent0108) Timestamp() time.Time { return time.Now() }
