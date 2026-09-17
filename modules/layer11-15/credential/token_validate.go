package credential

import (
	"time"
)

type TokenValidate struct{}

func NewTokenValidate() *TokenValidate {
	return &TokenValidate{}
}

func (t *TokenValidate) Validate() ([]SAMResult, error) {
	results := make([]SAMResult, 0, 1)
	results = append(results, SAMResult{Success: true, Method: SAMRegistryDump, Timestamp: time.Now()})
	return results, nil
}

func (t *TokenValidate) Name() string { return "TokenValidate" }
func (t *TokenValidate) Platform() string { return "windows" }
func (t *TokenValidate) RequiresElevation() bool { return false }
