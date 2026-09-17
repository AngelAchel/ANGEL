package osint

import (
	"time"
)

type osint0017 struct{}

func Newosint0017() *osint0017 {
	return &osint0017{}
}

func (e *osint0017) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "osint:done")
	return results, nil
}

func (e *osint0017) Name() string { return "osint0017" }
func (e *osint0017) Timestamp() time.Time { return time.Now() }
