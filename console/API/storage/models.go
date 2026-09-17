package storage

import (
	"time"
)

type Models struct{}

func NewModels() *Models {
	return &Models{}
}

func (m *Models) Get() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "models:done")
	return results, nil
}

func (m *Models) Name() string         { return "Models" }
func (m *Models) Timestamp() time.Time { return time.Now() }
