package osint

import (
	"time"
)

type osint0088 struct{}

func Newosint0088() *osint0088 {
	return &osint0088{}
}

func (e *osint0088) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "osint:done")
	return results, nil
}

func (e *osint0088) Name() string { return "osint0088" }
func (e *osint0088) Timestamp() time.Time { return time.Now() }
