package implant

import (
	"time"
)

type ImplantSleep struct{}

func NewImplantSleep() *ImplantSleep {
	return &ImplantSleep{}
}

func (i *ImplantSleep) Sleep() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "implant_sleep:done")
	return results, nil
}

func (i *ImplantSleep) Name() string         { return "ImplantSleep" }
func (i *ImplantSleep) Timestamp() time.Time { return time.Now() }
