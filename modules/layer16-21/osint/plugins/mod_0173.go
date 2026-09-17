package osint

import (
	"time"
)

type osint0173 struct{}

func Newosint0173() *osint0173 {
	return &osint0173{}
}

func (e *osint0173) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "osint:done")
	return results, nil
}

func (e *osint0173) Name() string { return "osint0173" }
func (e *osint0173) Timestamp() time.Time { return time.Now() }
