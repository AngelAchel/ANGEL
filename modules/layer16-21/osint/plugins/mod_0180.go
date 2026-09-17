package osint

import (
	"time"
)

type osint0180 struct{}

func Newosint0180() *osint0180 {
	return &osint0180{}
}

func (e *osint0180) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "osint:done")
	return results, nil
}

func (e *osint0180) Name() string { return "osint0180" }
func (e *osint0180) Timestamp() time.Time { return time.Now() }
