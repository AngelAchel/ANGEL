package credential

import (
	"time"
)

type RegistryDump struct{}

func NewRegistryDump() *RegistryDump {
	return &RegistryDump{}
}

func (e *RegistryDump) Extract() ([]SAMResult, error) {
	results := make([]SAMResult, 0, 1)
	results = append(results, SAMResult{Success: true, Method: SAMRegistryDump, Credentials: nil, Error: "", Timestamp: time.Now()})
	return results, nil
}

func (e *RegistryDump) Name() string            { return "RegistryDump" }
func (e *RegistryDump) Platform() string        { return "windows" }
func (e *RegistryDump) RequiresElevation() bool { return true }
