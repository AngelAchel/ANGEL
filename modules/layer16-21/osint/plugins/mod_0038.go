package osint

import (
	"time"
)

type osint0038 struct{}

func Newosint0038() *osint0038 {
	return &osint0038{}
}

func (e *osint0038) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "osint:done")
	return results, nil
}

func (e *osint0038) Name() string { return "osint0038" }
func (e *osint0038) Timestamp() time.Time { return time.Now() }
