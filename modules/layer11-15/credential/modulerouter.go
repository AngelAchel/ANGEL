package credential

import (
	"time"
)

// ModuleRouter routes modules between layers
type ModuleRouter struct{}

func NewModuleRouter() *ModuleRouter {
	return &ModuleRouter{}
}

func (r *ModuleRouter) Route(topic string) ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "modulerouter:routed")
	return results, nil
}

func (r *ModuleRouter) Name() string { return "ModuleRouter" }
func (r *ModuleRouter) Timestamp() time.Time { return time.Now() }
