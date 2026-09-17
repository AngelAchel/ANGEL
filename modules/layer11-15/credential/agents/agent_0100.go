package credential

import (
	"time"
)

type CredentialAgent0100 struct{}

func NewCredentialAgent0100() *CredentialAgent0100 {
	return &CredentialAgent0100{}
}

func (e *CredentialAgent0100) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "credential:done")
	return results, nil
}

func (e *CredentialAgent0100) Name() string         { return "CredentialAgent0100" }
func (e *CredentialAgent0100) Timestamp() time.Time { return time.Now() }
