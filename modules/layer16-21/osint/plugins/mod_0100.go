package osint

import (
	"time"
)

type osint0100 struct{}

func Newosint0100() *osint0100 {
	return &osint0100{}
}

func (e *osint0100) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "osint:done")
	return results, nil
}

func (e *osint0100) Name() string { return "osint0100" }
func (e *osint0100) Timestamp() time.Time { return time.Now() }
