package osint

import (
	"time"
)

type osint0000 struct{}

func Newosint0000() *osint0000 {
	return &osint0000{}
}

func (e *osint0000) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "osint:done")
	return results, nil
}

func (e *osint0000) Name() string { return "osint0000" }
func (e *osint0000) Timestamp() time.Time { return time.Now() }
