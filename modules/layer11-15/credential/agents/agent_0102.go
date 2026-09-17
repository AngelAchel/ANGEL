package credential

import (
	"time"
)

type CredentialAgent0102 struct{}

func NewCredentialAgent0102() *CredentialAgent0102 {
	return &CredentialAgent0102{}
}

func (e *CredentialAgent0102) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "credential:done")
	return results, nil
}

func (e *CredentialAgent0102) Name() string { return "CredentialAgent0102" }
func (e *CredentialAgent0102) Timestamp() time.Time { return time.Now() }
