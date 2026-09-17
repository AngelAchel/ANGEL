package implant

import (
	"time"
)

type ImplantMain struct{}

func NewImplantMain() *ImplantMain {
	return &ImplantMain{}
}

func (i *ImplantMain) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "implant_main:done")
	return results, nil
}

func (i *ImplantMain) Name() string { return "ImplantMain" }
func (i *ImplantMain) Timestamp() time.Time { return time.Now() }
