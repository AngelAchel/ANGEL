package sqlinject

import (
	"time"
)

type ErrorBased struct{}

func NewErrorBased() *ErrorBased {
	return &ErrorBased{}
}

func (e *ErrorBased) Inject() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "error_based:done")
	return results, nil
}

func (e *ErrorBased) Name() string { return "ErrorBased" }
func (e *ErrorBased) Timestamp() time.Time { return time.Now() }
