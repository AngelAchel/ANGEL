package credential

import (
	"time"
)

type CertExtractor struct{}

func NewCertExtractor() *CertExtractor {
	return &CertExtractor{}
}

func (c *CertExtractor) Extract() ([]SAMResult, error) {
	results := make([]SAMResult, 0, 1)
	results = append(results, SAMResult{Success: true, Method: SAMRegistryDump, Timestamp: time.Now()})
	return results, nil
}

func (c *CertExtractor) Name() string            { return "CertExtractor" }
func (c *CertExtractor) Platform() string        { return "windows" }
func (c *CertExtractor) RequiresElevation() bool { return true }
