package credential

import (
	"time"
)

type CredentialAgent0101 struct{}

func NewCredentialAgent0101() *CredentialAgent0101 {
	return &CredentialAgent0101{}
}

func (e *CredentialAgent0101) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "credential:done")
	return results, nil
}

func (e *CredentialAgent0101) Name() string         { return "CredentialAgent0101" }
func (e *CredentialAgent0101) Timestamp() time.Time { return time.Now() }
