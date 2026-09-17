package osint

import (
	"time"
)

type osint0102 struct{}

func Newosint0102() *osint0102 {
	return &osint0102{}
}

func (e *osint0102) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "osint:done")
	return results, nil
}

func (e *osint0102) Name() string { return "osint0102" }
func (e *osint0102) Timestamp() time.Time { return time.Now() }
