package credential

import (
	"time"
)

type UserExtract struct{}

func NewUserExtract() *UserExtract {
	return &UserExtract{}
}

func (e *UserExtract) Extract() ([]SAMResult, error) {
	results := make([]SAMResult, 0, 1)
	results = append(results, SAMResult{Success: true, Method: SAMRegistryDump, Credentials: nil, Error: "", Timestamp: time.Now()})
	return results, nil
}

func (e *UserExtract) Name() string            { return "UserExtract" }
func (e *UserExtract) Platform() string        { return "linux" }
func (e *UserExtract) RequiresElevation() bool { return false }
