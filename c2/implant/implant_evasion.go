package implant

import (
	"time"
)

type ImplantEvasion struct{}

func NewImplantEvasion() *ImplantEvasion {
	return &ImplantEvasion{}
}

func (i *ImplantEvasion) Evade() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "implant_evasion:done")
	return results, nil
}

func (i *ImplantEvasion) Name() string         { return "ImplantEvasion" }
func (i *ImplantEvasion) Timestamp() time.Time { return time.Now() }
