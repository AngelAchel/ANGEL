package credential

import (
	"time"
)

type PGShadow struct{}

func NewPGShadow() *PGShadow {
	return &PGShadow{}
}

func (e *PGShadow) Extract() ([]SAMResult, error) {
	results := make([]SAMResult, 0, 1)
	results = append(results, SAMResult{Success: true, Method: SAMRegistryDump, Credentials: nil, Error: "", Timestamp: time.Now()})
	return results, nil
}

func (e *PGShadow) Name() string { return "PGShadow" }
func (e *PGShadow) Platform() string { return "linux" }
func (e *PGShadow) RequiresElevation() bool { return true }
