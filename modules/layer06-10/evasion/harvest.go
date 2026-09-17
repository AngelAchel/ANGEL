package evasion

import (
	"time"
)

type Harvest struct{}

func NewHarvest() *Harvest {
	return &Harvest{}
}

func (e *Harvest) Gather() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "harvest:done")
	return results, nil
}

func (e *Harvest) Name() string         { return "Harvest" }
func (e *Harvest) Timestamp() time.Time { return time.Now() }
