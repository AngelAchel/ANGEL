package credential

import (
	"time"
)

type CredentialAgent0115 struct{}

func NewCredentialAgent0115() *CredentialAgent0115 {
	return &CredentialAgent0115{}
}

func (e *CredentialAgent0115) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "credential:done")
	return results, nil
}

func (e *CredentialAgent0115) Name() string         { return "CredentialAgent0115" }
func (e *CredentialAgent0115) Timestamp() time.Time { return time.Now() }
