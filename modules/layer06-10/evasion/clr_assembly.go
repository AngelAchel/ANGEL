package evasion

import (
	"time"
)

type CLRAssembly struct{}

func NewCLRAssembly() *CLRAssembly {
	return &CLRAssembly{}
}

func (c *CLRAssembly) Load() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "clr_assembly:loaded")
	return results, nil
}

func (c *CLRAssembly) Name() string         { return "CLRAssembly" }
func (c *CLRAssembly) Timestamp() time.Time { return time.Now() }
