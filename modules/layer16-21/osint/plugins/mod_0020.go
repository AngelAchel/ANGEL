package osint

import (
	"time"
)

type osint0020 struct{}

func Newosint0020() *osint0020 {
	return &osint0020{}
}

func (e *osint0020) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "osint:done")
	return results, nil
}

func (e *osint0020) Name() string { return "osint0020" }
func (e *osint0020) Timestamp() time.Time { return time.Now() }
