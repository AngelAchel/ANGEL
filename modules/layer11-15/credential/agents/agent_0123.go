package credential

import (
	"time"
)

type CredentialAgent0123 struct{}

func NewCredentialAgent0123() *CredentialAgent0123 {
	return &CredentialAgent0123{}
}

func (e *CredentialAgent0123) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "credential:done")
	return results, nil
}

func (e *CredentialAgent0123) Name() string         { return "CredentialAgent0123" }
func (e *CredentialAgent0123) Timestamp() time.Time { return time.Now() }
