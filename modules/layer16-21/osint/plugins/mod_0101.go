package osint

import (
	"time"
)

type osint0101 struct{}

func Newosint0101() *osint0101 {
	return &osint0101{}
}

func (e *osint0101) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "osint:done")
	return results, nil
}

func (e *osint0101) Name() string { return "osint0101" }
func (e *osint0101) Timestamp() time.Time { return time.Now() }
