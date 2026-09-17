package c2server

import (
	"time"
)

type Harvester struct{}

func NewHarvester() *Harvester {
	return &Harvester{}
}

func (e *Harvester) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "harvester:done")
	return results, nil
}

func (e *Harvester) Name() string         { return "Harvester" }
func (e *Harvester) Timestamp() time.Time { return time.Now() }
