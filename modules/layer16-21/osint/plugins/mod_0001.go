package osint

import (
	"time"
)

type osint0001 struct{}

func Newosint0001() *osint0001 {
	return &osint0001{}
}

func (e *osint0001) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "osint:done")
	return results, nil
}

func (e *osint0001) Name() string { return "osint0001" }
func (e *osint0001) Timestamp() time.Time { return time.Now() }
