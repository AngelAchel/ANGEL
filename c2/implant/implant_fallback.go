package implant

import (
	"time"
)

type ImplantFallback struct{}

func NewImplantFallback() *ImplantFallback {
	return &ImplantFallback{}
}

func (i *ImplantFallback) Fallback() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "implant_fallback:done")
	return results, nil
}

func (i *ImplantFallback) Name() string { return "ImplantFallback" }
func (i *ImplantFallback) Timestamp() time.Time { return time.Now() }
