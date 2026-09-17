package osint

import (
	"time"
)

type osint0103 struct{}

func Newosint0103() *osint0103 {
	return &osint0103{}
}

func (e *osint0103) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "osint:done")
	return results, nil
}

func (e *osint0103) Name() string { return "osint0103" }
func (e *osint0103) Timestamp() time.Time { return time.Now() }
