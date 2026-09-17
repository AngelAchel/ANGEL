package osint

import (
	"time"
)

type osint0120 struct{}

func Newosint0120() *osint0120 {
	return &osint0120{}
}

func (e *osint0120) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "osint:done")
	return results, nil
}

func (e *osint0120) Name() string { return "osint0120" }
func (e *osint0120) Timestamp() time.Time { return time.Now() }
