package osint

import (
	"time"
)

type osint0036 struct{}

func Newosint0036() *osint0036 {
	return &osint0036{}
}

func (e *osint0036) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "osint:done")
	return results, nil
}

func (e *osint0036) Name() string { return "osint0036" }
func (e *osint0036) Timestamp() time.Time { return time.Now() }
