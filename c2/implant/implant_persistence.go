package implant

import (
	"time"
)

type ImplantPersistence struct{}

func NewImplantPersistence() *ImplantPersistence {
	return &ImplantPersistence{}
}

func (i *ImplantPersistence) Persist() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "implant_persistence:done")
	return results, nil
}

func (i *ImplantPersistence) Name() string         { return "ImplantPersistence" }
func (i *ImplantPersistence) Timestamp() time.Time { return time.Now() }
