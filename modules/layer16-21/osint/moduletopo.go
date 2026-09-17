package osint

import (
	"time"
)

// ModuleTopology manages module topology
type ModuleTopology struct{}

func NewModuleTopology() *ModuleTopology {
	return &ModuleTopology{}
}

func (t *ModuleTopology) Topology() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "moduletopo:topology")
	return results, nil
}

func (t *ModuleTopology) Name() string         { return "ModuleTopology" }
func (t *ModuleTopology) Timestamp() time.Time { return time.Now() }
