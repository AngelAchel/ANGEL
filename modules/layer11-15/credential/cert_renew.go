package credential

import (
	"time"
)

type CertRenewal struct{}

func NewCertRenewal() *CertRenewal {
	return &CertRenewal{}
}

func (c *CertRenewal) Check() ([]SAMResult, error) {
	results := make([]SAMResult, 0, 1)
	results = append(results, SAMResult{Success: true, Method: SAMRegistryDump, Timestamp: time.Now()})
	return results, nil
}

func (c *CertRenewal) Name() string { return "CertRenewal" }
func (c *CertRenewal) Platform() string { return "windows" }
func (c *CertRenewal) RequiresElevation() bool { return false }
