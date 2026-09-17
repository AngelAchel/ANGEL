package osint

import (
	"time"
)

type osint0108 struct{}

func Newosint0108() *osint0108 {
	return &osint0108{}
}

func (e *osint0108) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "osint:done")
	return results, nil
}

func (e *osint0108) Name() string { return "osint0108" }
func (e *osint0108) Timestamp() time.Time { return time.Now() }
