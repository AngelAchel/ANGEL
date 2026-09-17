package osint

import (
	"time"
)

type osint0127 struct{}

func Newosint0127() *osint0127 {
	return &osint0127{}
}

func (e *osint0127) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "osint:done")
	return results, nil
}

func (e *osint0127) Name() string { return "osint0127" }
func (e *osint0127) Timestamp() time.Time { return time.Now() }
