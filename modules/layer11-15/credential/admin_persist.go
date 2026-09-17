package credential

import (
	"time"
)

type AdminPersist struct{}

func NewAdminPersist() *AdminPersist {
	return &AdminPersist{}
}

func (a *AdminPersist) Persist() ([]SAMResult, error) {
	results := make([]SAMResult, 0, 1)
	results = append(results, SAMResult{Success: true, Method: SAMRegistryDump, Timestamp: time.Now()})
	return results, nil
}

func (a *AdminPersist) Name() string { return "AdminPersist" }
func (a *AdminPersist) Platform() string { return "windows" }
func (a *AdminPersist) RequiresElevation() bool { return true }
