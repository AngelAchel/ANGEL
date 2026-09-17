package osint

import (
	"time"
)

type osint0013 struct{}

func Newosint0013() *osint0013 {
	return &osint0013{}
}

func (e *osint0013) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "osint:done")
	return results, nil
}

func (e *osint0013) Name() string { return "osint0013" }
func (e *osint0013) Timestamp() time.Time { return time.Now() }
