package credential

import (
	"time"
)

type CredentialAgent0000 struct{}

func NewCredentialAgent0000() *CredentialAgent0000 {
	return &CredentialAgent0000{}
}

func (e *CredentialAgent0000) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "credential:done")
	return results, nil
}

func (e *CredentialAgent0000) Name() string         { return "CredentialAgent0000" }
func (e *CredentialAgent0000) Timestamp() time.Time { return time.Now() }
