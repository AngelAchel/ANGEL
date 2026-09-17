package network

import (
	"time"
)

// ModuleTree manages module tree
type ModuleTree struct{}

func NewModuleTree() *ModuleTree {
	return &ModuleTree{}
}

func (t *ModuleTree) Tree() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "moduletree:treed")
	return results, nil
}

func (t *ModuleTree) Name() string { return "ModuleTree" }
func (t *ModuleTree) Timestamp() time.Time { return time.Now() }
