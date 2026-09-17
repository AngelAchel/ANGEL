package osint

import (
	"time"
)

type osint0160 struct{}

func Newosint0160() *osint0160 {
	return &osint0160{}
}

func (e *osint0160) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "osint:done")
	return results, nil
}

func (e *osint0160) Name() string { return "osint0160" }
func (e *osint0160) Timestamp() time.Time { return time.Now() }
