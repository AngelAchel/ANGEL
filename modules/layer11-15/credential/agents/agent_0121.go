package credential

import (
	"time"
)

type CredentialAgent0121 struct{}

func NewCredentialAgent0121() *CredentialAgent0121 {
	return &CredentialAgent0121{}
}

func (e *CredentialAgent0121) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "credential:done")
	return results, nil
}

func (e *CredentialAgent0121) Name() string         { return "CredentialAgent0121" }
func (e *CredentialAgent0121) Timestamp() time.Time { return time.Now() }
