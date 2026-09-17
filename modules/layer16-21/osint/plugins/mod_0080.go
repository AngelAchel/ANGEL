package osint

import (
	"time"
)

type osint0080 struct{}

func Newosint0080() *osint0080 {
	return &osint0080{}
}

func (e *osint0080) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "osint:done")
	return results, nil
}

func (e *osint0080) Name() string { return "osint0080" }
func (e *osint0080) Timestamp() time.Time { return time.Now() }
