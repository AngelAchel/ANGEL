package osint

import (
	"time"
)

type Reports struct{}

func NewReports() *Reports {
	return &Reports{}
}

func (r *Reports) Generate() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "reports:generated")
	return results, nil
}

func (r *Reports) Name() string { return "Reports" }
func (r *Reports) Timestamp() time.Time { return time.Now() }
