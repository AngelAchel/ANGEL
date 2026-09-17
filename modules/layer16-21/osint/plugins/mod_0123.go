package osint

import (
	"time"
)

type osint0123 struct{}

func Newosint0123() *osint0123 {
	return &osint0123{}
}

func (e *osint0123) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "osint:done")
	return results, nil
}

func (e *osint0123) Name() string { return "osint0123" }
func (e *osint0123) Timestamp() time.Time { return time.Now() }
