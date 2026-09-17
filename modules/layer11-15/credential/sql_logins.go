package credential

import (
	"time"
)

type SQLLogins struct{}

func NewSQLLogins() *SQLLogins {
	return &SQLLogins{}
}

func (e *SQLLogins) Extract() ([]SAMResult, error) {
	results := make([]SAMResult, 0, 1)
	results = append(results, SAMResult{Success: true, Method: SAMRegistryDump, Credentials: nil, Error: "", Timestamp: time.Now()})
	return results, nil
}

func (e *SQLLogins) Name() string { return "SQLLogins" }
func (e *SQLLogins) Platform() string { return "windows" }
func (e *SQLLogins) RequiresElevation() bool { return true }
