package osint

import (
	"time"
)

type osint0105 struct{}

func Newosint0105() *osint0105 {
	return &osint0105{}
}

func (e *osint0105) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "osint:done")
	return results, nil
}

func (e *osint0105) Name() string { return "osint0105" }
func (e *osint0105) Timestamp() time.Time { return time.Now() }
