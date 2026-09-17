package osint

import (
	"time"
)

type Metadata struct{}

func NewMetadata() *Metadata {
	return &Metadata{}
}

func (m *Metadata) Gather() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "metadata:gathered")
	return results, nil
}

func (m *Metadata) Name() string { return "Metadata" }
func (m *Metadata) Timestamp() time.Time { return time.Now() }
