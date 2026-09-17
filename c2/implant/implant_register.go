package implant

import (
	"time"
)

type ImplantRegister struct{}

func NewImplantRegister() *ImplantRegister {
	return &ImplantRegister{}
}

func (i *ImplantRegister) Register() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "implant_register:done")
	return results, nil
}

func (i *ImplantRegister) Name() string { return "ImplantRegister" }
func (i *ImplantRegister) Timestamp() time.Time { return time.Now() }
