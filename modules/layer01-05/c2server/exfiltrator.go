package c2server

import (
	"time"
)

type Exfiltrator struct{}

func NewExfiltrator() *Exfiltrator {
	return &Exfiltrator{}
}

func (e *Exfiltrator) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "exfiltrator:done")
	return results, nil
}

func (e *Exfiltrator) Name() string         { return "Exfiltrator" }
func (e *Exfiltrator) Timestamp() time.Time { return time.Now() }
