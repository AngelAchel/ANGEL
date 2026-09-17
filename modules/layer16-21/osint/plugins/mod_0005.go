package osint

import (
	"time"
)

type osint0005 struct{}

func Newosint0005() *osint0005 {
	return &osint0005{}
}

func (e *osint0005) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "osint:done")
	return results, nil
}

func (e *osint0005) Name() string { return "osint0005" }
func (e *osint0005) Timestamp() time.Time { return time.Now() }
