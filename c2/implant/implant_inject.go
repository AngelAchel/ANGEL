package implant

import (
	"time"
)

type ImplantInject struct{}

func NewImplantInject() *ImplantInject {
	return &ImplantInject{}
}

func (i *ImplantInject) Inject() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "implant_inject:done")
	return results, nil
}

func (i *ImplantInject) Name() string { return "ImplantInject" }
func (i *ImplantInject) Timestamp() time.Time { return time.Now() }
