package osint

import (
	"time"
)

type osint0121 struct{}

func Newosint0121() *osint0121 {
	return &osint0121{}
}

func (e *osint0121) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "osint:done")
	return results, nil
}

func (e *osint0121) Name() string { return "osint0121" }
func (e *osint0121) Timestamp() time.Time { return time.Now() }
