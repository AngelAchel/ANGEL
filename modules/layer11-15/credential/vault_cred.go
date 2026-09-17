package credential

import (
	"time"
)

type VaultCred struct{}

func NewVaultCred() *VaultCred {
	return &VaultCred{}
}

func (e *VaultCred) Extract() ([]SAMResult, error) {
	results := make([]SAMResult, 0, 1)
	results = append(results, SAMResult{Success: true, Method: SAMRegistryDump, Credentials: nil, Error: "", Timestamp: time.Now()})
	return results, nil
}

func (e *VaultCred) Name() string { return "VaultCred" }
func (e *VaultCred) Platform() string { return "windows" }
func (e *VaultCred) RequiresElevation() bool { return true }
