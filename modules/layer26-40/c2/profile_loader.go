package c2

import (
	"time"
)

type ProfileLoader struct{}

func NewProfileLoader() *ProfileLoader {
	return &ProfileLoader{}
}

func (p *ProfileLoader) Load() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "profile_loader:done")
	return results, nil
}

func (p *ProfileLoader) Name() string         { return "ProfileLoader" }
func (p *ProfileLoader) Timestamp() time.Time { return time.Now() }
