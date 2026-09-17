package credential

import (
	"time"
)

type CharFunc struct{}

func NewCharFunc() *CharFunc {
	return &CharFunc{}
}

func (c *CharFunc) Process() ([]SAMResult, error) {
	results := make([]SAMResult, 0, 1)
	results = append(results, SAMResult{Success: true, Method: SAMRegistryDump, Timestamp: time.Now()})
	return results, nil
}

func (c *CharFunc) Name() string { return "CharFunc" }
func (c *CharFunc) Platform() string { return "windows" }
func (c *CharFunc) RequiresElevation() bool { return false }
